// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func parseFlags() *Config {
	var config Config

	flag.StringVar(&config.dir, "dir", ".", "Directory to serve files from")
	flag.IntVar(&config.port, "port", 8000, "Port to bind server to")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: dev [options]\n\n")
		fmt.Fprintf(os.Stderr, "A simple static file server for local development.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if absDir, err := filepath.Abs(config.dir); err == nil {
		config.dir = absDir
	}

	return &config
}
