// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	version "github.com/caarlos0/go-version"
	"github.com/hugginsio/dev/handler"
	"github.com/hugginsio/dev/port"
)

type Config struct {
	dir  string
	port int
}

func main() {
	config := parseFlags()
	slog.Info("dev is starting", "dir", config.dir, "port", config.port, "version", version.GetVersionInfo().GitVersion)

	if _, err := os.Stat(config.dir); os.IsNotExist(err) {
		slog.Error("directory does not exist", "err", err.Error())
		os.Exit(http.StatusBadRequest)
	}

	port, err := port.Acquire(config.port)
	if err != nil {
		slog.Error("failed to find available port", "err", err.Error())
		os.Exit(http.StatusInternalServerError)
	}

	if config.port != port {
		slog.Warn("port unavailable, switching", "port", port)
	}

	handler := &handler.ServeHandler{Directory: config.dir}
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
