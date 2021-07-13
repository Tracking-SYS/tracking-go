package services

import (
	"context"

	entities_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/entities"

	"factory/exam/repo"
)

//ProductServiceInterface ...
type ProductServiceInterface interface {
	GetProducts(ctx context.Context, limit int) []*repo.ProductModel
	Transform(input []*repo.ProductModel) []*entities_pb.ProductInfo
}
