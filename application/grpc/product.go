package grpc

import (
	"context"
	"fmt"

	"github.com/igorlage/curso/fullcycle/desafio1/application/grpc/pb"
	"github.com/igorlage/curso/fullcycle/desafio1/application/usecase"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := p.ProductUseCase.RegisterProduct(in.Name, in.Description, float64(in.Price))
	if err != nil {
		return nil, fmt.Errorf("error while creating product: %v", err)
	}
	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
		},
	}, nil
}

func (p *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products, err := p.ProductUseCase.FindProducts()
	if err != nil {
		return &pb.FindProductsResponse{}, err
	}
	var responseProducts []*pb.Product
	for _, product := range products {
		responseProducts = append(responseProducts, &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
		})
	}
	return &pb.FindProductsResponse{
		Products: responseProducts,
	}, nil
}

func (p *ProductGrpcService) FindProduct(ctx context.Context, in *pb.FindProductRequest) (*pb.FindProductResponse, error) {
	product, err := p.ProductUseCase.FindProductById(in.Id)
	if err != nil {
		return nil, fmt.Errorf("error while fetching product: %v", err)
	}
	return &pb.FindProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       float32(product.Price),
	}, nil
}

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}
