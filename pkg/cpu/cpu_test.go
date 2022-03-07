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
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCheckProcessLiveness tests the process liveness check for the init process of PID 1
func TestCheckProcessLiveness(t *testing.T) {

	err := CheckProcessLiveness(1)
	assert.Nil(t, err, "Error occured during init process liveness check, %v", err)
}
