package rpc

import (
	"context"
	"io"
	"net/rpc"

	"github.com/fluxcd/flux/pkg/api/v12"
	"github.com/fluxcd/flux/pkg/remote"
)

// RPCClientV12 is the rpc-backed implementation of a server, for
// talking to remote daemons. This version introduces methods which accept an
// options struct as the first argument. e.g. ListServicesWithOptions
type RPCClientV12 struct {
	*RPCClientV11
}

type clientV12 interface {
	v12.Server
}

var _ clientV12 = &RPCClientV12{}

// NewClientV12 creates a new rpc-backed implementation of the server.
func NewClientV12(conn io.ReadWriteCloser) *RPCClientV12 {
	return &RPCClientV12{NewClientV11(conn)}
}

func (p *RPCClientV12) GitRepoConfigWithError(ctx context.Context, regenerate bool) (v12.GitConfig, error) {
	var result v12.GitConfig
	err := p.client.Call("RPCServer.GitRepoConfigWithError", regenerate, &result)
	if _, ok := err.(rpc.ServerError); !ok && err != nil {
		return v12.GitConfig{}, remote.FatalError{err}
	}
	if err != nil {
		err = remoteApplicationError(err)
	}
	return result, err
}
