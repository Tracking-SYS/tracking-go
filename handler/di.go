package handler

import (
	"github.com/google/wire"

	services_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/services"
)

//GraphSet ...
var GraphSet = wire.NewSet(
	NewProductPBHandler,
	wire.Bind(new(services_pb.ProductServiceServer), new(*ProductPBHandler)),

	NewTaskPBHandler,
	wire.Bind(new(services_pb.TaskServiceServer), new(*TaskPBHandler)),
)
