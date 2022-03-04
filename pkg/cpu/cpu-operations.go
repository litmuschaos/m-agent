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

// RevertStressNGProcess checks and reverts the defunct (zombie) stress-ng process
func RevertStressNGProcess(cmd *exec.Cmd, stderr *bytes.Buffer) error {

	if err := cmd.Wait(); err != nil {
		return errors.Errorf("stress-ng process exited with a non-zero exit code: %d; stderr: %v", cmd.ProcessState.ExitCode(), stderr.String())
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
	if err := CheckStressNGProcess(cmd.Process.Pid); err == nil {
		cmd.Wait()
	}

	return nil
}
