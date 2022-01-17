// Copyright 2022 LitmusChaos Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"time"
	"unicode"

	"github.com/denisbrodbeck/machineid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

// GenerateJWT generates a JWT for the authentication of client requests
func GenerateJWT(dayHourMinuteChar rune, duration int) (string, error) {

	// create a new token with the specified encryption algorithm
	token := jwt.New(jwt.SigningMethodHS256)

	// create a map for embedding any claims, if any, to be embedded into the token
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true

	// expire the token after this time
	switch unicode.ToLower(dayHourMinuteChar) {
	case 'd':
		claims["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(duration)).Unix()

	case 'h':
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(duration)).Unix()

	case 'm':
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(duration)).Unix()
	}

	machineId, err := machineid.ID()
	if err != nil {
		return "", errors.Errorf("failed to fetch the machine id, %v", err)
	}

	// sign the token with the secret signing key
	tokenString, err := token.SignedString([]byte(machineId))
	if err != nil {
		return "", errors.Errorf("failed to sign the authentication token, %v", err)
	}

	return tokenString, nil
}
