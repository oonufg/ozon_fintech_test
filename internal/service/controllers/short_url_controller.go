package server

import (
	"context"
	"ozon_fintech_test/internal/domain"
	"ozon_fintech_test/internal/persistence"
	pb "ozon_fintech_test/internal/service/proto/generated"
)

type ShortUrlController struct {
	repository persistence.UrlRepository
	pb.UnimplementedShortURLServer
}

func MakeShortUrlController(repository persistence.UrlRepository) *ShortUrlController {
	return &ShortUrlController{
		repository: repository,
	}
}

func (controller *ShortUrlController) CreateCompressedUrl(ctx context.Context, request *pb.CreateCompressedUrlRequest) (*pb.CreateCompressedUrlResponse, error) {
	compressedUrl := domain.CompressURL(request.FullUrl)
	error := controller.repository.SaveCompressedUrl(ctx, compressedUrl, request.FullUrl)
	if error != nil {
		return nil, error
	}
	return &pb.CreateCompressedUrlResponse{
		CompressedUrl: compressedUrl,
	}, nil
}

func (controller *ShortUrlController) GetFullUrlByCompressed(ctx context.Context, request *pb.GetFullUrlByCompressedRequest) (*pb.GetFullUrlByCompressedResponse, error) {
	result, error := controller.repository.GetFullUrlByCompressed(ctx, request.CompressedUrl)

	if error != nil {
		return nil, error
	}

	return &pb.GetFullUrlByCompressedResponse{
		FullUrl: result,
	}, nil
}
