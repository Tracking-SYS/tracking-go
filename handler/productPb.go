package handler

import (
	"context"

	"github.com/Tracking-SYS/tracking-go/services"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	servicesPb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/services"
)

//ProductPBHandler
type ProductPBHandler struct {
	servicesPb.UnimplementedProductServiceServer
	productService services.ProductServiceInterface
	tracer         trace.Tracer
}

//NewProductPBHandler
func NewProductPBHandler(
	productService services.ProductServiceInterface,
) *ProductPBHandler {
	tracer := otel.Tracer("NewProductPBHandler")
	return &ProductPBHandler{
		tracer:         tracer,
		productService: productService,
	}
}

//Get
func (p *ProductPBHandler) Get(ctx context.Context,
	req *servicesPb.ProductServiceGetRequest) (*servicesPb.ProductServiceGetResponse, error) {
	ctx, span := p.tracer.Start(ctx, "Get")
	defer span.End()

	limit := req.GetLimit()
	page := req.GetPage()
	ids := req.GetIds()
	products := p.productService.GetProducts(ctx, int(limit), int(page), ids)
	data := p.productService.Transform(products)

	return &servicesPb.ProductServiceGetResponse{
		Data: data,
	}, nil
}

//GetSingle
func (p *ProductPBHandler) GetSingle(ctx context.Context,
	req *servicesPb.ProductServiceGetSingleRequest) (*servicesPb.ProductServiceGetSingleResponse, error) {
	ctx, span := p.tracer.Start(ctx, "GetSingle")
	defer span.End()

	ID := req.GetId()
	product := p.productService.GetProduct(ctx, int(ID))
	data := p.productService.TransformSingle(product)

	return &servicesPb.ProductServiceGetSingleResponse{
		Data: data,
	}, nil
}

//Create
func (p *ProductPBHandler) Create(ctx context.Context,
	req *servicesPb.ProductServiceCreateRequest) (*servicesPb.ProductServiceCreateResponse, error) {
	ctx, span := p.tracer.Start(ctx, "Create")
	defer span.End()

	data := req.GetData()
	product := p.productService.CreateProduct(ctx, data)
	data = p.productService.TransformSingle(product)

	return &servicesPb.ProductServiceCreateResponse{
		Data: data,
	}, nil
}
