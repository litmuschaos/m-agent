package cpu

import (
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
