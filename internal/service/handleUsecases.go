package server

import (
	"context"
	"ozon_fintech_test/internal/domain"
	pb "ozon_fintech_test/internal/service/proto/generated"
)

func (server *Server) CreateCompressedUrl(ctx context.Context, request *pb.CreateCompressedUrlRequest) *pb.CreateCompressedUrlResponse {
	compressedUrl := domain.CompressURL(request.FullUrl)
	error := server.repository.SaveCompressedUrl(ctx, compressedUrl, request.FullUrl)
	if error != nil {
		return nil
	}
	return &pb.CreateCompressedUrlResponse{
		CompressedUrl: compressedUrl,
	}
}

func (server *Server) GetFullUrlByCompressed(ctx context.Context, request *pb.GetFullUrlByCompressedRequest) *pb.GetFullUrlByCompressedResponse {
	result, _ := server.repository.GetFullUrlByCompressed(ctx, request.CompressedUrl)
	return &pb.GetFullUrlByCompressedResponse{
		FullUrl: result,
	}
}
