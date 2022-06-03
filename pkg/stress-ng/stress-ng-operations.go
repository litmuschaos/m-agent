package stressng

import (
	"bytes"
	"os"
	"os/exec"
	"syscall"

	"github.com/pkg/errors"
)

// CheckStressNG uses a bash command to check if the stress-ng tool is installed
func CheckStressNG() error {

	cmd := exec.Command("/bin/sh", "-c", "command -v stress-ng")

	if err := cmd.Run(); err != nil {
		return errors.Errorf("stress-ng not found")
	}

	return nil
}

// RevertStressNGProcess checks and reverts the defunct (zombie) stress-ng process
func RevertStressNGProcess(cmd *exec.Cmd, stderr *bytes.Buffer) error {

	if err := cmd.Wait(); err != nil {
		return errors.Errorf("stress-ng process exited with a non-zero exit code: %d; stderr: %v", cmd.ProcessState.ExitCode(), stderr.String())
	}

	return nil
}

// CheckStressNGProcessLiveness checks if a given process is currently running
func CheckStressNGProcessLiveness(pid int) error {

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

// AbortStressNGProcess kills a running stress-ng process, and if the
// process has already completed, it clears the defunct (zombie) process
func AbortStressNGProcess(cmd *exec.Cmd) error {

	// kill the running stress-ng process to make it exit immediately
	if err := cmd.Process.Kill(); err != nil {
		return errors.Errorf("failed to kill the stress-ng process, err: %v", err)
	}

	// kill will not be able to exit a defunct (zombie) process,
	// which will be present only if the stress-ng process
	// has already completed. Hence if the process isn't killed,
	// we wait on it, which immediately clears the defunct (zombie) process
	if err := CheckStressNGProcessLiveness(cmd.Process.Pid); err == nil {
		cmd.Wait()
	}

	return nil
}
