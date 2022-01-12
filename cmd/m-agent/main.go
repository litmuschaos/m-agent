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

	"github.com/litmuschaos/m-agent/api/server"
	"github.com/litmuschaos/m-agent/internal/m-agent/port"
	"github.com/litmuschaos/m-agent/internal/m-agent/tokens"
)

func main() {

	generateToken := flag.Bool("get-token", false, "generates a token to be used for the authentication of the requests made to the agent")
	tokenExpiryDuration := flag.String("token-expiry-duration", "", "token expiry duration (non-interactive mode)")
	flag.Parse()

	if *generateToken {

		// generate a JWT for authentication
		if *tokenExpiryDuration == "" {

			tokens.HandleInteractiveTokenGeneration()
		} else {

			tokens.HandleNonInteractiveTokenGeneration(*tokenExpiryDuration)
		}
	} else {
		if port.IsPortOpen("41365") {

			// handle client requests
			server.HandleRequests()
		} else {

			// port is not open i.e. daemon is already running
			flag.Usage()
		}
	}
}
