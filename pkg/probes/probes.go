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

package probes

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
)

// ExecuteCmdProbeCommand executes a shell command sent by the client as part of the cmd probe validation
func ExecuteCmdProbeCommand(payload []byte) (string, error) {

	var command string

	if err := json.Unmarshal(payload, &command); err != nil {
		return "", err
	}

	var stdout, stderr bytes.Buffer

	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", errors.New(err.Error() + "; error output: " + stderr.String())
	}

	return stdout.String(), nil
}
