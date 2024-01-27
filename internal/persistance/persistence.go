package persistence

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UrlRepository interface {
	GetFullUrlByCompressed(ctx context.Context, url string) (string, error)
	SaveCompressedUrl(ctx context.Context, string, compressedUrl string) error
}

type InmemoryUrlRepository struct {
	UrlRepository
	table map[string]string
}

type PostgreSQLUrlRepository struct {
	UrlRepository
	database *pgxpool.Pool
}

func MakeInmemoryRepository() *InmemoryUrlRepository {
	return &InmemoryUrlRepository{
		table: make(map[string]string),
	}
}

func MakePGRepository(db *pgxpool.Pool) *PostgreSQLUrlRepository {
	return &PostgreSQLUrlRepository{
		database: db,
	}
}

func (imr *InmemoryUrlRepository) GetFullUrlByCompressed(ctx context.Context, url string) (string, error) {
	value, isExist := imr.table[url]
	if isExist == false {
		return "", errors.New("Such url not found")
	}
	return value, nil
}

func (imr *InmemoryUrlRepository) SaveCompressedUrl(ctx context.Context, fullUrl string, compressedUrl string) error {
	imr.table[compressedUrl] = fullUrl
	return nil
}

func (pgr *PostgreSQLUrlRepository) GetFullUrlByCompressed(ctx context.Context, compressedUrl string) (string, error) {
	result, error := GetFullUrlByCompressedFromSQL(ctx, pgr.database, compressedUrl)
	if error != nil {
		return "", error
	}
	return result, nil
}

func (pgr *PostgreSQLUrlRepository) SaveCompressedUrl(ctx context.Context, compressedUrl string, fulUrl string) {
	SaveURLInSQLDB(ctx, pgr.database, compressedUrl, fulUrl)
}
