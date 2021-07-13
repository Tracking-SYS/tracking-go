package handler

import (
	"github.com/google/wire"

	services_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/services"
)

//GraphSet ...
var GraphSet = wire.NewSet(
	ProductHandlerProvider,

	NewProductPBHandler,
	wire.Bind(new(services_pb.ProductServiceServer), new(*ProductPBHandler)),
)
