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
	"log"

	"github.com/gorilla/websocket"
)

type Message struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"body"`
	ReqID   string      `json:"reqid"`
}

// ListenForClientMessage listens for client messages and returns the received message action and payload
func ListenForClientMessage(conn *websocket.Conn) (string, string, []byte, error) {

	var msg Message

	if err := conn.ReadJSON(&msg); err != nil {
		return "", "", []byte{}, err
	}

	payload, err := json.Marshal(msg.Payload)
	if err != nil {
		return "", "", []byte{}, err
	}

	return msg.Action, msg.ReqID, payload, nil
}

// SendMessageToClient wraps a message action and payload in a Message structure and sends it to the client
func SendMessageToClient(conn *websocket.Conn, action, reqID string, payload interface{}) error {

	return conn.WriteJSON(Message{action, payload, reqID})
}

// HandleActionExecutionError handles an error generated while performing the action requested by the client
func HandleActionExecutionError(conn *websocket.Conn, reqID, errorCode string, err error, logger *log.Logger) {

	if err := SendMessageToClient(conn, "ERROR", reqID, errorCode+err.Error()); err != nil {
		logger.Printf("Error occured while sending error message to client, err: %v", err)
	}

	if err := conn.Close(); err != nil {
		logger.Printf("Error occured while closing the connection, err: %v", err)
	}
}

// HandleFeedbackTransmissionError handles an error generated while sending the feedback for an earlier requested action to the client
func HandleFeedbackTransmissionError(conn *websocket.Conn, err error, logger *log.Logger) {

	logger.Printf("Error occured while sending feedback message to client, err: %v", err)

	if err := conn.Close(); err != nil {
		logger.Printf("Error occured while closing the connection, err: %v", err)
	}
}
