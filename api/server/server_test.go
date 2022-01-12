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

package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/litmuschaos/m-agent/api/server/auth"
	processKill "github.com/litmuschaos/m-agent/experiments/process-kill"
	errorcodes "github.com/litmuschaos/m-agent/internal/m-agent/error-codes"
	"github.com/stretchr/testify/assert"
)

// TestFallbackRouteHandler tests the fallback route
func TestFallbackRouteHandler(t *testing.T) {

	testCases := []string{"/", "/abc", "/xyz/a", "/process", "string", "with-hyphen", "with@special@characters", "12345"}

	for _, testCase := range testCases {
		req, err := http.NewRequest(http.MethodGet, testCase, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(fallbackRouteHandler)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code, "Expected status code: %v, but received %v", http.StatusNotFound, rr.Code)

		expectedErrorMessage := "Invalid Access: requested route not found"

		assert.Equal(t, expectedErrorMessage, rr.Body.String(), "Unexpected error message: %s", rr.Body.String())
	}
}

// TestProcessKillEndpointNoToken tests the middleware for no token in request header
func TestProcessKillEndpointNoToken(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "/process-kill", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(auth.IsAuthorized(processKill.ProcessKill))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected status code: %v, but received %v", http.StatusBadRequest, rr.Code)

	expectedErrorMessage := errorcodes.GetAuthErrorPrefix() + "Invalid Access: authentication token not found"

	assert.Equal(t, expectedErrorMessage, rr.Body.String())
}

// TestProcessKillEndpointInvalidToken tests the middleware for an invalid token in request header
func TestProcessKillEndpointInvalidToken(t *testing.T) {

	testCases := []string{"abc", "/abc", "1234", "with@special@characters", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjbGllbnQiOiJOZWVsYW5qYW4iLCJleHAiOjE2MzczMjk0MTF9.ys2quDrxR0eR2A5nfvSU0_OiuLnPaQaoEuo8dW-VSi4", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjbGllbnQiOiJOZWVsYW5qYW4iLCJleHAiOjE2MzczMjk0MTF9"}

	for _, testCase := range testCases {

		req, err := http.NewRequest(http.MethodGet, "/process-kill", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header = http.Header{"Authorization": []string{"Bearer " + testCase}}

		rr := httptest.NewRecorder()
		handler := http.Handler(auth.IsAuthorized(processKill.ProcessKill))
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code, "Expected status code: %v, but received %v", http.StatusUnauthorized, rr.Code)
	}
}
