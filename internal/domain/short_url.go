package domain

import (
	"crypto/sha256"
	"encoding/hex"
)

func CompressURL(link string) string {
	result := sha256.Sum256([]byte(link))
	return hex.EncodeToString(result[:])[:10]
}
