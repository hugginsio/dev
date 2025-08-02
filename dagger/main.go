// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"dagger/dev/internal/dagger"
)

type Dev struct{}

// Build a ready-to-use development environment.
func (m *Dev) devEnv(
	ctx context.Context,
	source *dagger.Directory,
	// +optional
	platform *dagger.Platform,
) *dagger.Container {
	if platform == nil {
		enginePlatform, err := dag.DefaultPlatform(ctx)
		if err != nil {
			panic(err)
		}

		platform = &enginePlatform
	}

	return dag.Container(dagger.ContainerOpts{Platform: *platform}).
		From("golang:1.24-alpine").
		WithExec([]string{"apk", "add", "--no-cache", "git"}).
		WithDirectory("/go/src/", source).
		WithMountedCache("/go/pkg/mod/", dag.CacheVolume("go-mod-124")).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", dag.CacheVolume("go-build-124")).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithWorkdir("/go/src/").
		WithExec([]string{"go", "mod", "download"})
}
