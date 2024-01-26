package server

import (
	"ozon_fintech_test/internal/domain"
)

func (server *Server) CreateCompressedUrl(fullUrl string) string {
	compressedUrl := domain.CompressURL(fullUrl)
	error := server.repository.SaveCompressedUrl(compressedUrl, fullUrl)
	if error != nil {
		return ""
	}
	return compressedUrl
}

func (server *Server) GetFullUrlByCompressed(compressedUrl string) string {
	result, _ := server.repository.GetFullUrlByCompressed(compressedUrl)
	return result

}
