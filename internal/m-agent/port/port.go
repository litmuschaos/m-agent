package port

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

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

// GetMAgentPort returns the m-agent server port from the PORT config file
func GetMAgentPort() (string, error) {

	serverPortSlice, err := os.ReadFile("/etc/m-agent/PORT")
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(serverPortSlice), "\n"), nil
}

// UpdateMAgentPort updates the m-agent server port by updating the PORT config file
func UpdateMAgentPort(newPort string) error {

	if !IsPortValid(newPort) {
		return errors.Errorf("invalid port")
	}

	if !IsPortOpen(newPort) {
		return errors.Errorf("port unavailable")
	}

	f, err := os.OpenFile("/etc/eg/PORT", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer f.Close()

	if err := f.Truncate(0); err != nil {
		return err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, "%s", newPort+"\n")
	if err != nil {
		return err
	}

	return nil
}
