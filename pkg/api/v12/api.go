// This package defines the types for Flux API version 12.
package v12

import (
	"context"

	v11 "github.com/fluxcd/flux/pkg/api/v11"
	v6 "github.com/fluxcd/flux/pkg/api/v6"
)

// GitConfig extends the v6.GitConfig with Error
type GitConfig struct {
	v6.GitConfig
	Error string `json:"errors"`
}

// Server in version 12
type Server interface {
	v11.Server

	GitRepoConfigWithError(ctx context.Context, regenerate bool) (GitConfig, error)
}
