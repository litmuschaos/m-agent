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
	"os"
	"os/exec"
	"syscall"

	"github.com/pkg/errors"
)

// CheckStressNG uses a bash command to check if the stress-ng tool is installed
func CheckStressNG() error {

	var stderr bytes.Buffer

	cmd := exec.Command("/bin/sh", "-c", "command stress-ng")
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.Errorf(stderr.String())
	}

	return nil
}

// CheckProcessLiveness checks if a given process is currently running
func CheckProcessLiveness(pid int) error {

	// On Unix systems, FindProcess always succeeds and returns a Process for
	// the given pid, regardless of whether the process exists.
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	// If 0 signal is assigned to be sent to the process, then no signal is sent, but error checking is still performed;
	// this can be used to check for the existence of a process ID
	if err := p.Signal(syscall.Signal(0)); err != nil {
		return errors.Errorf("received error on sending 0 signal to the stress-ng process, err: %v", err)
	}

	return nil
}
