package repo

import (
	"context"

	entities_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/entities"
)

//ProductRepoInterface ...
type ProductRepoInterface interface {
	Get(context context.Context, limit int, page int, ids []uint64) (productDAO []*ProductModel, err error)
	Find(context context.Context, id int) (productDAO *ProductModel, err error)
	Create(context context.Context, id *entities_pb.ProductInfo) (productDAO *ProductModel, err error)
}
