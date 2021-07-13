package handler

import (
	"context"
	"factory/exam/services"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	services_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/services"
)

//ProductPBHandler ...
type ProductPBHandler struct {
	services_pb.UnimplementedProductServiceServer
	productService services.ProductServiceInterface
	tracer         trace.Tracer
}

//NewProductPBHandler ...
func NewProductPBHandler(
	productService services.ProductServiceInterface,
) *ProductPBHandler {
	tracer := otel.Tracer("ProductHandlerProvider")
	return &ProductPBHandler{
		tracer:         tracer,
		productService: productService,
	}
}

//Get ...
func (p *ProductPBHandler) Get(ctx context.Context,
	req *services_pb.ProductServiceGetRequest) (*services_pb.ProductServiceGetResponse, error) {
	ctx, span := p.tracer.Start(ctx, "Get")
	defer span.End()

	limit := req.GetLimit()
	products := p.productService.GetProducts(ctx, int(limit))
	data := p.productService.Transform(products)

	return &services_pb.ProductServiceGetResponse{
		Data: data,
	}, nil
}
