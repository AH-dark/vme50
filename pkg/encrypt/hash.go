package encrypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/AH-dark/random-donate/pkg/conf"
)

func HmacSha256(msg string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}

func Pass(pass string) string {
	return HmacSha256(pass, conf.SystemConfig.HashSecret)
}
