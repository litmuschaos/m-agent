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

package process

import (
	"encoding/json"
	"testing"
)

// TestProcessStateCheck executes the Process State Check with a single process of PID 1
func TestProcessStateCheck(t *testing.T) {

	pids := []int{1}

	payload, err := json.Marshal(pids)
	if err != nil {
		t.Fatalf("Error occured while marshalling PIDs, %v", err)
	}

	if err := ProcessStateCheck(payload); err != nil {
		t.Fatalf("Error occured during process state check, %v", err)
	}
}
