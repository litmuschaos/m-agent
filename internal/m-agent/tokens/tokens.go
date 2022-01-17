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
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/fatih/color"
	"github.com/litmuschaos/m-agent/api/server/auth"
	errorcodes "github.com/litmuschaos/m-agent/internal/m-agent/error-codes"
	"github.com/litmuschaos/m-agent/internal/m-agent/ip-address"
	"github.com/manifoldco/promptui"
)

// HandleInteractiveTokenGeneration facilitates the generation of a JWT with an expiry time with an interactive CLI prompt
func HandleInteractiveTokenGeneration() {

	// set token error code and token error string
	log.SetPrefix(errorcodes.GetTokenErrorPrefix())

	var token string
	var err error

	tokenExpirationPrompts := []string{"30 Minutes", "1 Hour", "24 Hours", "30 Days"}

	list := promptui.Select{
		Label: "The token should expire after",
		Items: tokenExpirationPrompts,
	}

	idx, _, err := list.Run()
	if err != nil {
		log.Printf("Error during token expiry prompt selection, %v", err)
		return
	}

	switch tokenExpirationPrompts[idx] {
	case "30 Minutes":
		token, err = auth.GenerateJWT('m', 30)
		if err != nil {
			log.Printf("Error during authentication token generation, %v", err)
			return
		}
	case "1 Hour":
		token, err = auth.GenerateJWT('h', 1)
		if err != nil {
			log.Printf("Error during authentication token generation, %v", err)
			return
		}
	case "24 Hours":
		token, err = auth.GenerateJWT('h', 24)
		if err != nil {
			log.Printf("Error during authentication token generation, %v", err)
			return
		}
	case "30 Days":
		token, err = auth.GenerateJWT('d', 30)
		if err != nil {
			log.Printf("Error during authentication token generation, %v", err)
			return
		}
	}

	endpoint := ip.GetPublicIP() + ":41365"

	boldWhite := color.New(color.FgWhite, color.Bold)

	boldWhite.Print("Agent Endpoint: ")
	fmt.Println(endpoint)

	boldWhite.Print("Authentication Token: ")
	fmt.Println(token)
}

// HandleNonInteractiveTokenGeneration facilitates the generation of a JWT with an expiry time in a non-interactive manner
func HandleNonInteractiveTokenGeneration(tokenExpiryDuration string) {

	// set token error code and token error string
	log.SetPrefix(errorcodes.GetTokenErrorPrefix())

	type Token struct {
		Token    string `json:"token"`
		Endpoint string `json:"endpoint"`
	}

	dayHourMinuteChar, dayHourMinuteValue, err := validateTokenExpiryDuration(tokenExpiryDuration)
	if err != nil {
		log.Println("Invalid token expiry duration")
		return
	}

	token, err := auth.GenerateJWT(dayHourMinuteChar, dayHourMinuteValue)
	if err != nil {
		log.Printf("Error during authentication token generation, %v", err)
		return
	}

	endpoint := ip.GetPublicIP() + ":41365"

	jsonResult, err := json.MarshalIndent(Token{Token: token, Endpoint: endpoint}, "", "  ")
	if err != nil {
		log.Printf("Error during creation of JSON token output, %v", err)
		return
	}

	jsonResultString := string(jsonResult)

	// json marshalling replaces '<' and '>' characters with their unicode value
	// hence the unicode values are replaced back with the original characters in the resultant string
	jsonResultString = strings.Replace(jsonResultString, "\\u003c", "<", -1)
	jsonResultString = strings.Replace(jsonResultString, "\\u003e", ">", -1)

	fmt.Println(jsonResultString)
}

// validateTokenExpiryDuration ensures that the token expiry duration input is correctly formatted
func validateTokenExpiryDuration(tokenExpiryDuration string) (rune, int, error) {

	// fetch the last character in the tokenExpiryDuration string
	dayHourMinuteChar := []rune(tokenExpiryDuration[len(tokenExpiryDuration)-1:])[0]

	if !unicode.IsLetter(dayHourMinuteChar) {
		return ' ', 0, errors.New("")
	}

	// slice-off the last character in the tokenExpiryDuration string
	dayHourMinuteValue, err := strconv.Atoi(tokenExpiryDuration[:len(tokenExpiryDuration)-1])
	if err != nil {
		return ' ', 0, err
	}

	switch unicode.ToLower(dayHourMinuteChar) {
	case 'd':
		if dayHourMinuteValue < 1 || dayHourMinuteValue > 30 {
			return ' ', 0, errors.New("")
		}

	case 'h':
		if dayHourMinuteValue < 1 || dayHourMinuteValue > 24 {
			return ' ', 0, errors.New("")
		}

	case 'm':
		if dayHourMinuteValue < 1 || dayHourMinuteValue > 60 {
			return ' ', 0, errors.New("")
		}

	default:
		return ' ', 0, errors.New("")
	}

	return dayHourMinuteChar, dayHourMinuteValue, nil
}
