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
	"fmt"
	"net/http"
	"strings"

	"github.com/denisbrodbeck/machineid"
	"github.com/golang-jwt/jwt/v4"
	errorcodes "github.com/litmuschaos/m-agent/internal/m-agent/errorcodes"
	"github.com/pkg/errors"
)

// IsAuthorized validates whether the client request is authenticated with a valid JWT
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Authorization") != "" {

			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]

			// parse the JWT token using the secret signing key, after verifying the ecryption algorithm
			token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {

				// verify that the token is encrypted with HS256 ecryption algorithm
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("cannot verify HS256 ecryption")
				}

				machineId, err := machineid.ID()
				if err != nil {
					return nil, errors.Errorf("failed to fetch the machine id, %v", err)
				}

				return []byte(machineId), nil
			})

			if err != nil {

				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, errorcodes.GetAuthErrorPrefix()+"Invalid Authentication: Failed to authenticate using the token, "+err.Error())
				return
			}

			// redirect to the intended route handler
			if token.Valid {
				endpoint(w, r)
			}
		} else {

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, errorcodes.GetAuthErrorPrefix()+"Invalid Access: authentication token not found")
			return
		}
	})
}
