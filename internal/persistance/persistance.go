package persistence

import (
	"database/sql"
	"errors"
)

type UrlRepository interface {
	GetFullUrlByCompressed(url string) (string, error)
	SaveCompressedUrl(url string, compressedUrl string) error
}

type InmemoryUrlRepository struct {
	UrlRepository
	table map[string]string
}

type PostgreSQLUrlRepository struct {
	UrlRepository
	database *sql.DB
}

func MakeInmemoryRepository() *InmemoryUrlRepository {
	return &InmemoryUrlRepository{
		table: make(map[string]string),
	}
}

func (imr *InmemoryUrlRepository) GetFullUrlByCompressed(url string) (string, error) {
	value, isExist := imr.table[url]
	if isExist == false {
		return "", errors.New("Such url not found")
	}
	return value, nil
}

func (imr *InmemoryUrlRepository) SaveCompressedUrl(fullUrl string, compressedUrl string) error {
	imr.table[compressedUrl] = fullUrl
	return nil
}

func (pgr *PostgreSQLUrlRepository) GetFullUrlByCompressed(compressedUrl string) (string, error) {
	result, error := GetFullUrlByCompressedFromSQL(pgr.database, compressedUrl)
	if error != nil {
		return "", error
	}
	return result, nil
}

func (pgr *PostgreSQLUrlRepository) SaveCompressedUrl(compressedUrl string, fulUrl string) {
	SaveURLInSQLDB(pgr.database, compressedUrl, fulUrl)
}
