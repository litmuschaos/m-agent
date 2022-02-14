package cpu

import (
	"bytes"
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

// CheckForStressNGProcess checks if the stress-ng process has successfully exited or not
// If the process is still running then it forcefully kills the process and returns
func CheckForStressNGProcess(cmd *exec.Cmd, stderr bytes.Buffer) error {

	if !cmd.ProcessState.Exited() {

		if err := syscall.Kill(cmd.Process.Pid, 9); err != nil {
			return errors.Errorf("failed to force stop the stress-ng process, err: %v", err)
		}

		return nil
	}

	if !cmd.ProcessState.Success() {
		return errors.Errorf("stress-ng process failed during execution with %v exit code, err: %v", cmd.ProcessState.ExitCode(), stderr.String())
	}

	return nil
}
