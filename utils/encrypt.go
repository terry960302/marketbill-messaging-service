package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func HMAC256(payload string, secret string) string {
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(payload))
	return base64.StdEncoding.EncodeToString(hmac256.Sum(nil))
}
