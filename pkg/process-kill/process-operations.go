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
	"syscall"
)

// killProcess kills a given target process
func killProcess(processId, signal int) error {

	return syscall.Kill(processId, syscall.Signal(signal))
}

// KillTargetProcesses kills all the target processes
func KillTargetProcesses(payload []byte) error {

	var processes []int

	if err := json.Unmarshal(payload, &processes); err != nil {
		return err
	}

	for _, processId := range processes {

		if err := killProcess(processId, 9); err != nil {
			return err
		}
	}

	return nil
}
