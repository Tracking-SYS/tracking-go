package services

import (
	"context"
	"factory/exam/repo"

	entities_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/entities"
)

var _ ProductServiceInterface = &ProductService{}

//ProductProvider ...
func ProductProvider(
	productRepo repo.ProductRepoInterface,
) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

//ProductService ...
type ProductService struct {
	productRepo repo.ProductRepoInterface
}

//GetProducts ...
func (ps *ProductService) GetProducts(ctx context.Context, limit int) []*repo.ProductModel {
	products, err := ps.productRepo.GetProduct(ctx, limit)
	if err != nil {
		return nil
	}

	return products
}

//Transform ...
func (ps *ProductService) Transform(input []*repo.ProductModel) []*entities_pb.ProductInfo {
	result := []*entities_pb.ProductInfo{}
	for _, prod := range input {
		item := &entities_pb.ProductInfo{
			Id:    uint32(prod.ID),
			Name:  prod.Name,
			Price: prod.Price,
			Type:  prod.Type,
		}
		result = append(result, item)
	}

	return result
}
