// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import "log/slog"

type Config struct {
	dir  string
	port int
}

func main() {
	config := parseFlags()
	slog.Info("dev is starting", "dir", config.dir, "port", config.port)
}
