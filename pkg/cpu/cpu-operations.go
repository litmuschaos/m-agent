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

package cpu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

// StressCPU starts a stress-ng process in background and returns the exec cmd for it
func StressCPU(payload []byte, stdout, stderr *bytes.Buffer) (*exec.Cmd, error) {

	type CPUStressParams struct {
		Workers string
		Load    string
		Timeout string
	}

	var cpuStressParams CPUStressParams

	if err := json.Unmarshal(payload, &cpuStressParams); err != nil {
		return nil, err
	}

	stressCommand := fmt.Sprintf("stress-ng --cpu %s --cpu-load %s --timeout %s", cpuStressParams.Workers, cpuStressParams.Load, cpuStressParams.Timeout)

	cmd := exec.Command("bash", "-c", stressCommand)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		return nil, errors.Errorf("%s, stderr: %s", err, stderr.String())
	}

	return cmd, nil
}
