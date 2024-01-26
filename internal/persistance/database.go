package persistence

import (
	"database/sql"
)

const (
	SAVE_URL = `INSERT INTO short_urls(compressed_url, full_url) VALUES($1, $2)`

	GET_FULL_URL_BY_COMPRESSED = `
		SELECT full_url FROM short_urls
		WHERE compressed_url = $1
		`
)

func SaveURLInSQLDB(database *sql.DB, compressedURL string, fullUrl string) error {
	transaction, error := database.Begin()
	if error != nil {
		return error
	}

	defer func() {
		if panicValue := recover(); panicValue != nil {
			transaction.Rollback()
		}
	}()

	_, error = transaction.Exec(SAVE_URL, compressedURL, fullUrl)
	if error != nil {
		transaction.Rollback()
	}

	error = transaction.Commit()
	if error != nil {
		return error
	}

	return nil
}

func GetFullUrlByCompressedFromSQL(database *sql.DB, compressedURL string) (string, error) {
	transaction, error := database.Begin()
	if error != nil {
		return "", error
	}
	var fullUrl string

	error = transaction.QueryRow(GET_FULL_URL_BY_COMPRESSED, compressedURL).Scan(&fullUrl)
	if error != nil {
		return "", error
	}

	return fullUrl, nil
}
