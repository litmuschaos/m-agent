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

package messages

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Message struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"body"`
}

// ListenForClientMessage listens for client messages and returns the received message action and payload
func ListenForClientMessage(conn *websocket.Conn) (string, []byte, error) {

	var msg Message

	if err := conn.ReadJSON(&msg); err != nil {
		return "", []byte{}, err
	}

	payload, err := json.Marshal(msg.Payload)
	if err != nil {
		return "", []byte{}, err
	}

	return msg.Action, payload, nil
}

// SendMessageToClient wraps a message action and payload in a Message structure and sends it to the client
func SendMessageToClient(conn *websocket.Conn, action string, payload interface{}) error {

	return conn.WriteJSON(Message{action, payload})
}
