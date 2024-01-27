package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	SAVE_URL = `INSERT INTO short_urls(compressed_url, full_url) VALUES($1, $2)`

	GET_FULL_URL_BY_COMPRESSED = `
		SELECT full_url FROM short_urls
		WHERE compressed_url = $1
		`
)

func SaveURLInSQLDB(ctx context.Context, database *pgxpool.Pool, compressedURL string, fullUrl string) error {
	transaction, error := database.Begin(ctx)
	if error != nil {
		return error
	}

	defer func() {
		if panicValue := recover(); panicValue != nil {
			transaction.Rollback(ctx)
		}
	}()

	_, error = transaction.Exec(ctx, SAVE_URL, compressedURL, fullUrl)
	if error != nil {
		transaction.Rollback(ctx)
	}

	error = transaction.Commit(ctx)
	if error != nil {
		return error
	}

	return nil
}

func GetFullUrlByCompressedFromSQL(ctx context.Context, database *pgxpool.Pool, compressedURL string) (string, error) {
	transaction, error := database.Begin(ctx)
	if error != nil {
		return "", error
	}
	var fullUrl string

	error = transaction.QueryRow(ctx, GET_FULL_URL_BY_COMPRESSED, compressedURL).Scan(&fullUrl)
	if error != nil {
		return "", error
	}

	return fullUrl, nil
}
