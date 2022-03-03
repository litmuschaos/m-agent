package cpu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/gorilla/websocket"

	"github.com/pkg/errors"
)

// StressCPU starts a stress-ng process in background and returns the exec cmd for it
func StressCPU(payload []byte, reqID string, stdout, stderr *bytes.Buffer, conn *websocket.Conn) (*exec.Cmd, error) {

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

// RevertStressNGProcess checks the stress-ng process exit code and cleans up the defunct process
func RevertStressNGProcess(cmd *exec.Cmd, stderr *bytes.Buffer) error {

	if err := cmd.Wait(); err != nil {
		return errors.Errorf("stress-ng process exited with a non-zero exit code %d, stderr: %v", cmd.ProcessState.ExitCode(), stderr.String())
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
