package cpu

import (
	"bytes"
	"os"
	"os/exec"
	"syscall"

	"github.com/pkg/errors"
)

// CheckForStressNG uses a bash command to check if the stress-ng tool is installed
func CheckForStressNG() error {

	cmd := exec.Command("/bin/sh", "-c", "command -v stress-ng")

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// AbortStressNGProcess checks if the stress-ng process has successfully exited or not.
// If the process is still running then it forcefully kills the process and returns
func AbortStressNGProcess(cmd *exec.Cmd) error {

	if !cmd.ProcessState.Exited() {

		if err := syscall.Kill(cmd.Process.Pid, 9); err != nil {
			return errors.Errorf("failed to force stop the stress-ng process, err: %v", err)
		}

		return nil
	}

	return nil
}

// CheckStressNGProcessLiveness checks if the stress-ng process is running
func CheckStressNGProcessLiveness(cmd *exec.Cmd, stderr *bytes.Buffer) error {

	p, err := os.FindProcess(cmd.Process.Pid)
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
