package main

import (
	"context"
	"dagger/dev/internal/dagger"
)

// Build the executable binary
func (m *Dev) Build(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.File {
	// TODO: multi arch
	return m.devEnv(ctx, source, nil).
		WithWorkdir("cmd/dev").
		WithEnvVariable("CGO_ENABLED", "0").
		WithExec([]string{"go", "build", "-ldflags", "-s -w", "-gcflags=all=-l -C", "-buildvcs", "-o", "/app/dev", "."}).
		File("/app/dev")
}

// Releases the dev CLI with goreleaser
func (m *Dev) ReleaseCli(
	ctx context.Context,
	tag string,
	// +optional
	token *dagger.Secret,
) (string, error) {
	source := dag.Git("https://github.com/hugginsio/dev.git", dagger.GitOpts{KeepGitDir: true}).Tag(tag).Tree()

	return dag.Container().
		From("ghcr.io/goreleaser/goreleaser:v2.11.2").
		WithSecretVariable("GITHUB_TOKEN", token).
		WithMountedCache("/go/pkg/mod/", dag.CacheVolume("go-mod-124")).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", dag.CacheVolume("go-build-124")).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithDirectory("/go/src/github.com/hugginsio/dev/", source).
		WithWorkdir("/go/src/github.com/hugginsio/dev/").
		WithExec([]string{"goreleaser", "release"}).
		Stdout(ctx)
}
