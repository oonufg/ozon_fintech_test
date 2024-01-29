\c short_urls_service;

CREATE TABLE IF NOT EXISTS short_urls(
    compressed_url VARCHAR(10) NOT NULL,
    full_url VARCHAR NOT NULL
)
