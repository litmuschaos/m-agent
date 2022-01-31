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

package main

import (
	"flag"
	"log"

	"github.com/litmuschaos/m-agent/api/server"
	logger "github.com/litmuschaos/m-agent/internal/m-agent/log"
	"github.com/litmuschaos/m-agent/internal/m-agent/port"
	"github.com/litmuschaos/m-agent/internal/m-agent/tokens"
)

func main() {

	generateToken := flag.Bool("get-token", false, "generates a token to be used for the authentication of the requests made to the agent")
	tokenExpiryDuration := flag.String("token-expiry-duration", "", "token expiry duration (non-interactive mode)")
	serverPort := flag.String("port", "", "port for m-agent")
	flag.Parse()

	if *generateToken {

		// set token error code and token error string
		tokenErrorLogger := logger.GetTokenErrorLogger()

		// generate a JWT for authentication
		if *tokenExpiryDuration == "" {

			if err := tokens.HandleInteractiveTokenGeneration(); err != nil {
				tokenErrorLogger.Println(err)
			}
		} else {

			if err := tokens.HandleNonInteractiveTokenGeneration(*tokenExpiryDuration); err != nil {
				tokenErrorLogger.Println(err)
			}
		}
	} else {

		// if port.IsPortOpen("41365") {

		// 	// handle client requests
		// 	if err := server.HandleRequests(); err != nil {
		// 		log.Fatal(err)
		// 	}
		// } else {

		// 	// port is not open i.e. daemon is already running
		// 	flag.Usage()
		// }

		if *serverPort != "" {
			// maybe not necessary
			if !port.IsPortValid(*serverPort) {
				log.Fatalf("%v port is invalid", *serverPort)
				return
			}

			if !port.IsPortOpen(*serverPort) {
				log.Fatalf("%v port is not available", *serverPort)
				return
			}

			if err := server.HandleRequests(*serverPort); err != nil {
				log.Fatal(err)
			}
		} else {

			flag.Usage()
		}
	}
}
