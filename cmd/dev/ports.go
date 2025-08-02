// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"fmt"
	"net"
	"strconv"
)

func acquirePort(startPort int) (int, error) {
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
