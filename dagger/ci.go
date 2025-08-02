// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"dagger/dev/internal/dagger"
)

// Lint the project.
func (m *Dev) Lint(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	return m.devEnv(ctx, source, nil).WithExec([]string{"go", "vet", "./..."}).Stdout(ctx)
}
