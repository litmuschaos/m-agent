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

package experiment

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/litmuschaos/m-agent/internal/m-agent/messages"
	"github.com/stretchr/testify/assert"
)

// sendMessageToServer sends messages to the client by encapsulating an action and a payload
func sendMessageToServer(conn *websocket.Conn, action string, payload interface{}) (string, []byte, error) {

	if err := conn.WriteJSON(messages.Message{Action: action, Payload: payload}); err != nil {
		return "", nil, err
	}

	var msg messages.Message

	if err := conn.ReadJSON(&msg); err != nil {
		return "", nil, err
	}

	clientPayload, err := json.Marshal(msg.Payload)
	if err != nil {
		return "", nil, err
	}

	return msg.Action, clientPayload, nil
}

// TestProcessKillCmdProbe executes the cmd probe
func TestProcessKillCmdProbe(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(ProcessKill))

	// Convert http://127.0.0.1 to ws://127.0.0.1
	url := "ws" + strings.TrimPrefix(s.URL, "http")

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	defer func() {
		conn.Close()
		s.Close()
	}()

	feedback, payload, err := sendMessageToServer(conn, "EXECUTE_COMMAND", `echo "This is echoed from the probe"`)
	if err != nil {
		t.Fatalf("Unable to send message to the server, %v", err)
	}

	if feedback != "ACTION_SUCCESSFUL" {

		var serverError string

		if feedback == "ERROR" {

			if err := json.Unmarshal(payload, &serverError); err != nil {
				t.Fatalf("Failed to interpret error message from server, %v", err)
			}

			t.Fatalf(serverError)
		}

		t.Fatalf("Unintelligible feedback: %v", feedback)
	}

	var stdout string

	if err := json.Unmarshal(payload, &stdout); err != nil {
		t.Fatalf("Failed to interpret message from server, %v", err)
	}

	stdout = strings.TrimSpace(stdout)

	expectedOutput := "This is echoed from the probe"

	assert.Equal(t, expectedOutput, stdout, "Expected String: %s", expectedOutput)
}

func TestProcessKillSteadyStateCheck(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(ProcessKill))

	// Convert http://127.0.0.1 to ws://127.0.0.1
	url := "ws" + strings.TrimPrefix(s.URL, "http")

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	defer func() {
		conn.Close()
		s.Close()
	}()

	feedback, payload, err := sendMessageToServer(conn, "CHECK_STEADY_STATE", []int{1})
	if err != nil {
		t.Fatalf("unable to send message to the server, %v", err)
	}

	if feedback != "ACTION_SUCCESSFUL" {

		var serverError string

		if feedback == "ERROR" {

			if err := json.Unmarshal(payload, &serverError); err != nil {
				t.Fatalf("Failed to interpret error message from server, %v", err)
			}

			t.Fatalf(serverError)
		}

		t.Fatalf("Unintelligible feedback: %v", feedback)
	}
}
