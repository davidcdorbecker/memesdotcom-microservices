package helpers

import (
	"crypto/sha256"
	"encoding/base64"
)

func Encrypt(s string) string {
	h := sha256.Sum256([]byte(s))
	return base64.StdEncoding.EncodeToString(h[:])
}
