package server

import (
	"github.com/google/wire"

	"factory/exam/config"
	"factory/exam/handler"
	"factory/exam/infra"
	repo_bind "factory/exam/repo/bind"
	kafka_server "factory/exam/server/kafka"
	"factory/exam/services"
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
	kafka_server.NewKafkaConsumer,
	kafka_server.NewKafkaProducer,
)
