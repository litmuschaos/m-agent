package port

import "net"

// IsPortOpen returns if a port is open or not
func IsPortOpen(port string) bool {

	ln, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return false
	}

	_ = ln.Close()

	return true
}
