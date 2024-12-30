package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Hex(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
