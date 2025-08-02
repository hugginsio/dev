// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type Config struct {
	dir  string
	port int
}

func main() {
	config := parseFlags()
	slog.Info("dev is starting", "dir", config.dir, "port", config.port)

	if _, err := os.Stat(config.dir); os.IsNotExist(err) {
		slog.Error("directory does not exist", "err", err.Error())
		os.Exit(http.StatusBadRequest)
	}

	port, err := acquirePort(config.port)
	if err != nil {
		slog.Error("failed to find available port", "err", err.Error())
		os.Exit(http.StatusInternalServerError)
	}

	if config.port != port {
		slog.Warn("port unavailable, switching", "port", port)
	}

	handler := &devHandler{dir: config.dir}
	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	slog.Info("server is up", "url", fmt.Sprintf("http://localhost:%d", port))

	if err := server.ListenAndServe(); err != nil {
		slog.Error("failure", "err", err.Error())
		os.Exit(http.StatusInternalServerError)
	}
}
