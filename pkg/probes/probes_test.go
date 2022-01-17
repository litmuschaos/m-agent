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
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestExecuteCmdProbeCommand executes the cmdProbe with an echo command
func TestExecuteCmdProbeCommand(t *testing.T) {

	command := `echo "This is printed from the probe"`
	expectedOutput := "This is printed from the probe"

	payload, err := json.Marshal(command)
	if err != nil {
		t.Fatal(err)
	}

	stdout, err := ExecuteCmdProbeCommand(payload)
	if err != nil {
		t.Fatal(err)
	}

	stdout = strings.TrimSpace(stdout)

	assert.Equal(t, stdout, expectedOutput, "Expected the stdout to be, '%s'", expectedOutput)
}
