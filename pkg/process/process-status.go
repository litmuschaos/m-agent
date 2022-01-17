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

	"github.com/mitchellh/go-ps"
	"github.com/pkg/errors"
)

// isProcessRunning checks if a given process is running or not
func isProcessRunning(processId int) (bool, error) {

	p, err := ps.FindProcess(processId)
	if err != nil {
		return false, err
	}

	if p != nil {
		return true, nil
	}

	return false, nil
}

// ProcessStateCheck validates that all the target processes are running
func ProcessStateCheck(payload []byte) error {

	// var processes Processes
	var processes []int

	if err := json.Unmarshal(payload, &processes); err != nil {
		return err
	}

	if len(processes) == 0 {
		return errors.New("no process found")
	}

	for _, processId := range processes {

		isProcessRunning, err := isProcessRunning(processId)
		if err != nil {
			return err
		}

		if !isProcessRunning {
			return errors.Errorf("%v process not found", processId)
		}
	}

	return nil
}
