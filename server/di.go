package server

import (
	"github.com/google/wire"

	"github.com/Tracking-SYS/tracking-go/config"
	"github.com/Tracking-SYS/tracking-go/handler"
	"github.com/Tracking-SYS/tracking-go/infra"
	repo_bind "github.com/Tracking-SYS/tracking-go/repo/bind"
	"github.com/Tracking-SYS/tracking-go/services"
)

//ServerDeps ...
var ServerDeps = wire.NewSet(
	config.GraphSet,
	handler.GraphSet,
	services.GraphSet,
	repo_bind.GraphSet,
	infra.GraphSet,
)

//GraphSet ...
var GraphSet = wire.NewSet(
	ServerDeps,
	HTTPProvider,
	NewMetricServer,
	NewServerManager,
	NewKafkaConsumer,
)
