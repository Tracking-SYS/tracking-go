// +build wireinject

package main

import (
	"context"

	"github.com/Tracking-SYS/tracking-go/server"

	"github.com/google/wire"
)

func buildServer(ctx context.Context) (*server.Manager, error) {
	panic(wire.Build(server.GraphSet))
}
