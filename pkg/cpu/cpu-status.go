package cpu

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/pkg/errors"
)

// CheckStressNG uses a bash command to check if the stress-ng tool is installed
func CheckStressNG() error {

	cmd := exec.Command("/bin/sh", "-c", "command -v stress-ng")

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// CheckStressNGProcess checks if the stress-ng process is running
func CheckStressNGProcess(pid int) error {

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
