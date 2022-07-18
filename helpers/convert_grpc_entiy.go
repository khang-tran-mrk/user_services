package helpers

import (
	"ChoTot/entity"
	product_grpc "ChoTot/grpc"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func P_GrpcToEntity(product *product_grpc.CreateProductRequest) *entity.Product {
	return &entity.Product{
		ProductName: product.ProductName,
		UserId:      1,
		CatId:       product.CatId,
		TypeId:      product.TypeId,
		Price:       product.Price,
		State:       product.State,
		//CreatedTime: product.CreatedTime.AsTime(),
		//ExpiredTime: product.ExpiredTime.AsTime(),
		Address: product.Address,
		Content: product.Content,
	}
}

func P_EntityToGrpc(product *entity.Product) *product_grpc.ProductEntity {
	return &product_grpc.ProductEntity{
		Id:          int32(product.Id),
		ProductName: product.ProductName,
		CatId:       product.CatId,
		TypeId:      product.TypeId,
		Price:       product.Price,
		State:       product.State,
		CreatedTime: timestamppb.New(product.CreatedTime),
		ExpiredTime: timestamppb.New(product.ExpiredTime),
		Priority:    product.Priority,
		Address:     product.Address,
		Content:     product.Content,
		Images:      nil,
	}
}

func P_EntityToGrpcCreateReq(product *entity.Product) *product_grpc.CreateProductRequest {
	return &product_grpc.CreateProductRequest{
		ProductName: product.ProductName,
		CatId:       product.CatId,
		TypeId:      product.TypeId,
		Price:       product.Price,
		State:       product.State,
		Address:     product.Address,
		Content:     product.Content,
	}
}
