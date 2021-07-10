package services

import (
	"context"
	"encoding/json"
	"factory/exam/repo"
	"factory/exam/repo/cache"
	"fmt"
	"strconv"

	entities_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/entities"
)

var _ ProductServiceInterface = &ProductService{}

//ProductProvider ...
func ProductProvider(
	productRepo repo.ProductRepoInterface,
	cacheRepo cache.CacheInteface,
) *ProductService {
	return &ProductService{
		productRepo: productRepo,
		cacheRepo:   cacheRepo,
	}
}

//ProductService ...
type ProductService struct {
	productRepo repo.ProductRepoInterface
	cacheRepo   cache.CacheInteface
}

//GetProducts ...
func (ps *ProductService) GetProducts(ctx context.Context, limit int, page int, ids []uint64) []*repo.ProductModel {
	products, err := ps.productRepo.Get(ctx, limit, page, ids)
	if err != nil {
		return nil
	}

	return products
}

//GetProduct ...
func (ps *ProductService) GetProduct(ctx context.Context, id int) *repo.ProductModel {
	product, err := ps.cacheRepo.Get(ctx, strconv.Itoa(id))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if product != nil {
		fmt.Printf("GetCache: %v\n", product)
		return ps.parseData(product.(map[string]interface{}))
	}

	product, err = ps.productRepo.Find(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = ps.cacheRepo.Set(ctx, fmt.Sprintf("product_%s", strconv.Itoa(id)), product)
	if err != nil {
		fmt.Println(err)
	}

	return product.(*repo.ProductModel)
}

func (ps *ProductService) parseData(data map[string]interface{}) (product *repo.ProductModel) {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := json.Unmarshal(jsonbody, &product); err != nil {
		fmt.Println(err)
		return
	}

	return product
}

func (ps *ProductService) CreateProduct(ctx context.Context, data *entities_pb.ProductInfo) *repo.ProductModel {
	product, err := ps.productRepo.Create(ctx, data)
	if err != nil {
		return nil
	}

	return product
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

//Transform ...
func (ps *ProductService) TransformSingle(prod *repo.ProductModel) *entities_pb.ProductInfo {
	if prod == nil {
		return nil
	}

	item := &entities_pb.ProductInfo{
		Id:    uint32(prod.ID),
		Name:  prod.Name,
		Price: prod.Price,
		Type:  prod.Type,
	}

	return item
}
