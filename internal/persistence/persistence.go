package persistence

import (
	"context"
	"errors"
	"fmt"
	"log"
	"ozon_fintech_test/cfg"

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

func (imr *InmemoryUrlRepository) SaveCompressedUrl(ctx context.Context, compressedUrl string, fullUrl string) error {
	imr.table[compressedUrl] = fullUrl
	return nil
}

func (pgr *PostgreSQLUrlRepository) GetFullUrlByCompressed(ctx context.Context, compressedUrl string) (string, error) {
	result, error := GetFullUrlByCompressedFromSQL(ctx, pgr.database, compressedUrl)
	if error != nil {
		return "", errors.New("Such url not found")
	}
	return result, nil
}

func (pgr *PostgreSQLUrlRepository) SaveCompressedUrl(ctx context.Context, compressedUrl string, fulUrl string) error {
	SaveURLInSQLDB(ctx, pgr.database, compressedUrl, fulUrl)
	return nil
}

func RepositoryFactoryMethod(cfg *cfg.Cfg) UrlRepository {
	if cfg.PERSISTANCE_MODE == "inmemory" {
		return MakeInmemoryRepository()
	} else if cfg.PERSISTANCE_MODE == "database" {
		dbConPool, error := pgxpool.New(context.TODO(), fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.DATABASE_USER, cfg.DATABASE_PASSWORD, cfg.DATABASE_HOST, cfg.DATABASE_PORT, cfg.DATABASE))
		if error != nil {
			log.Fatal("Failed connect to database")
		}
		return MakePGRepository(dbConPool)
	} else {
		log.Println("Type of storage not recognized, supports: 'inmemory', 'database'")
		log.Println("Chosen Default: 'inmemory'")
		return MakeInmemoryRepository()
	}
}
