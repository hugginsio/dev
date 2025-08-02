// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Port provides a utility for local port acquisition.
package port

import (
	"fmt"
	"net"
	"strconv"
)

// Acquire attempts to bind a local port. If the specified port is unavailable, it will retry until
// either a port is found or port 65535 is reached. If no available ports are found, an error will be returned.
func Acquire(startPort int) (int, error) {
	for port := startPort; port < 65535; port++ {
		addr := net.JoinHostPort("", strconv.Itoa(port))
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			listener.Close()
			return port, nil
		}
	}

	return 0, fmt.Errorf("no available ports found")
}
