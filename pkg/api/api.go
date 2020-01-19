package api

import v12 "github.com/fluxcd/flux/pkg/api/v12"

// Server defines the minimal interface a Flux must satisfy to adequately serve a
// connecting fluxctl. This interface specifically does not facilitate connecting
// to Weave Cloud.
type Server interface {
	v12.Server
}
