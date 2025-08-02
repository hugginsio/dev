package main

import (
	"context"
	"dagger/dev/internal/dagger"
)

// Run all pull request checks.
func (m *Dev) PullRequest(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	// TODO: traces
	lint, err := m.Lint(ctx, source)
	if err != nil {
		return "", err
	}

	return lint, nil
}
