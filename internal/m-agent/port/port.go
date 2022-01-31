package port

import (
	"net"
	"strconv"

	"github.com/mitchellh/go-ps"
	"github.com/pkg/errors"
)

//IsPortValid checks if a port value is valid or not
func IsPortValid(port string) bool {

	portValue, err := strconv.Atoi(port)
	if err != nil {
		return false
	}

	if portValue < 1 || portValue > 65535 || port[0] == '0' {
		return false
	}

	return true
}

// IsPortOpen checks if a port is open or not
func IsPortOpen(port string) bool {

	ln, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return false
	}

	_ = ln.Close()

	return true
}

// GetMAgentPort returns the port at which m-agent is running. If the process is not running then an empty string is returned
func GetMAgentPort() (string, error) {

	isProcess, err := isMAgentProcessRunning()
	if err != nil {
		return "", errors.Errorf("failed to check for m-agent process, err: %v", err)
	}

	if !isProcess {
		return "", nil
	}

	return "[PORT]", nil
}

func isMAgentProcessRunning() (bool, error) {

	processes, err := ps.Processes()
	if err != nil {
		return false, err
	}

	for _, p := range processes {

		if p.Executable() == "m-agent" {
			return true, nil
		}
	}

	return false, nil
}
