package cpu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

// StressCpu starts a stress-ng process in background and returns the exec cmd for it
func StressCpu(payload []byte, stdout, stderr *bytes.Buffer) (*exec.Cmd, error) {

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
