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
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/litmuschaos/m-agent/api/server/auth"
	cpuStress "github.com/litmuschaos/m-agent/experiments/cpu-stress/experiment"
	processKill "github.com/litmuschaos/m-agent/experiments/process-kill/experiment"
)

// fallbackRouteHandler serves a 404 status code and error message for invalid routes
func fallbackRouteHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Invalid Access: requested route not found")
}

// HandleRequests listens for requests made by the clients at any specified route
func HandleRequests() error {

	router := mux.NewRouter()

	router.Handle("/process-kill", auth.IsAuthorized(processKill.ProcessKill))
	router.Handle("/cpu-stress", auth.IsAuthorized(cpuStress.CPUStress))
	router.NotFoundHandler = http.HandlerFunc(fallbackRouteHandler)

	return http.ListenAndServe(":41365", router)
}
