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

package tokens

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/pkg/errors"

	"github.com/fatih/color"
	"github.com/litmuschaos/m-agent/api/server/auth"
	"github.com/litmuschaos/m-agent/internal/m-agent/ip"
	"github.com/manifoldco/promptui"
)

// HandleInteractiveTokenGeneration facilitates the generation of a JWT with an expiry time with an interactive CLI prompt
func HandleInteractiveTokenGeneration() error {

	var token string
	var err error

	tokenExpirationPrompts := []string{"30 Minutes", "1 Hour", "24 Hours", "30 Days"}

	list := promptui.Select{
		Label: "The token should expire after",
		Items: tokenExpirationPrompts,
	}

	idx, _, err := list.Run()
	if err != nil {
		return errors.Errorf("Error during token expiry prompt selection, %v", err)
	}

	switch tokenExpirationPrompts[idx] {
	case "30 Minutes":
		token, err = auth.GenerateJWT('m', 30)
		if err != nil {
			return errors.Errorf("Error during authentication token generation with 30 min validity, %v", err)
		}
	case "1 Hour":
		token, err = auth.GenerateJWT('h', 1)
		if err != nil {
			return errors.Errorf("Error during authentication token generation with 1 hr validity, %v", err)
		}
	case "24 Hours":
		token, err = auth.GenerateJWT('h', 24)
		if err != nil {
			return errors.Errorf("Error during authentication token generation with 24 hr validity, %v", err)
		}
	case "30 Days":
		token, err = auth.GenerateJWT('d', 30)
		if err != nil {
			return errors.Errorf("Error during authentication token generation with 30 days validity, %v", err)
		}
	}

	endpoint := ip.GetPublicIP() + ":41365"

	boldWhite := color.New(color.FgWhite, color.Bold)

	boldWhite.Print("Agent Endpoint: ")
	fmt.Println(endpoint)

	boldWhite.Print("Authentication Token: ")
	fmt.Println(token)

	return nil
}

// HandleNonInteractiveTokenGeneration facilitates the generation of a JWT with an expiry time in a non-interactive manner
func HandleNonInteractiveTokenGeneration(tokenExpiryDuration string) error {

	type Token struct {
		Endpoint string `json:"endpoint"`
		Token    string `json:"token"`
	}

	dayHourMinuteChar, duration, err := validateTokenExpiryDuration(tokenExpiryDuration)
	if err != nil {
		return err
	}

	token, err := auth.GenerateJWT(dayHourMinuteChar, duration)
	if err != nil {
		return errors.Errorf("Error during authentication token generation, %v", err)
	}

	endpoint := ip.GetPublicIP() + ":41365"

	jsonResult, err := json.MarshalIndent(Token{Endpoint: endpoint, Token: token}, "", "  ")
	if err != nil {
		return errors.Errorf("Error during creation of JSON token output, %v", err)
	}

	jsonResultString := string(jsonResult)

	// json marshalling replaces '<' and '>' characters with their unicode value
	// hence the unicode values are replaced back with the original characters in the resultant string
	jsonResultString = strings.Replace(jsonResultString, "\\u003c", "<", -1)
	jsonResultString = strings.Replace(jsonResultString, "\\u003e", ">", -1)

	fmt.Println(jsonResultString)

	return nil
}

// validateTokenExpiryDuration ensures that the token expiry duration input is correctly formatted
func validateTokenExpiryDuration(tokenExpiryDuration string) (rune, int, error) {

	// fetch the last character in the tokenExpiryDuration string
	dayHourMinuteChar := []rune(tokenExpiryDuration[len(tokenExpiryDuration)-1:])[0]

	if !unicode.IsLetter(dayHourMinuteChar) {
		return ' ', 0, errors.Errorf("Invalid token expiry duration")
	}

	// slice-off the last character in the tokenExpiryDuration string
	duration, err := strconv.Atoi(tokenExpiryDuration[:len(tokenExpiryDuration)-1])
	if err != nil {
		return ' ', 0, err
	}

	switch unicode.ToLower(dayHourMinuteChar) {
	case 'd':
		if duration < 1 || duration > 30 {
			return ' ', 0, errors.Errorf("Invalid token expiry duration")
		}

	case 'h':
		if duration < 1 || duration > 24 {
			return ' ', 0, errors.Errorf("Invalid token expiry duration")
		}

	case 'm':
		if duration < 1 || duration > 60 {
			return ' ', 0, errors.Errorf("Invalid token expiry duration")
		}

	default:
		return ' ', 0, errors.Errorf("Invalid token expiry duration")
	}

	return dayHourMinuteChar, duration, nil
}
