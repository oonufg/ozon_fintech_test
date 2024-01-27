package server

import (
	"context"
	"ozon_fintech_test/internal/domain"
	pb "ozon_fintech_test/internal/service/proto/generated"
)

func (server *Server) CreateCompressedUrl(ctx context.Context, request *pb.CreateCompressedUrlRequest) (*pb.CreateCompressedUrlResponse, error) {
	compressedUrl := domain.CompressURL(request.FullUrl)
	error := server.repository.SaveCompressedUrl(ctx, compressedUrl, request.FullUrl)
	if error != nil {
		return nil, error
	}
	return &pb.CreateCompressedUrlResponse{
		CompressedUrl: compressedUrl,
	}, nil
}

func (server *Server) GetFullUrlByCompressed(ctx context.Context, request *pb.GetFullUrlByCompressedRequest) (*pb.GetFullUrlByCompressedResponse, error) {
	result, error := server.repository.GetFullUrlByCompressed(ctx, request.CompressedUrl)

	if error != nil {
		return nil, error
	}

	return &pb.GetFullUrlByCompressedResponse{
		FullUrl: result,
	}, nil
}
