package services

import (
	"context"

	entities_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/entities"

	"github.com/Tracking-SYS/tracking-go/repo"
)

//ProductServiceInterface ...
type ProductServiceInterface interface {
	GetProducts(ctx context.Context, limit int, page int, ids []uint64) []*repo.ProductModel
	GetProduct(ctx context.Context, id int) *repo.ProductModel
	CreateProduct(ctx context.Context, data *entities_pb.ProductInfo) *repo.ProductModel
	Transform(input []*repo.ProductModel) []*entities_pb.ProductInfo
	TransformSingle(prod *repo.ProductModel) *entities_pb.ProductInfo
}
