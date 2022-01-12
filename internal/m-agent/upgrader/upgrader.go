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

package upgrader

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// GetConnectionUpgrader returns an upgrader to promote the connection to a bidirectional websocket
func GetConnectionUpgrader() websocket.Upgrader {

	// upgrader defines the read and write buffer size for the websocket
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true }, // allows CORS
	}

	return upgrader
}
