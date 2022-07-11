package encrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"github.com/AH-dark/random-donate/pkg/conf"
	"github.com/AH-dark/random-donate/pkg/utils"
	"io"
)

func HmacSha256(msg string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}

func HmacMD5(msg string, secret string) string {
	h := hmac.New(md5.New, []byte(secret))
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}

func MD5(msg string) string {
	m := md5.New()
	_, err := io.WriteString(m, msg)
	if err != nil {
		utils.Log().Error("error in md5 func: %v", err.Error())
		return ""
	}

	return hex.EncodeToString(m.Sum(nil))
}

func Pass(pass string) string {
	return HmacSha256(pass, conf.SystemConfig.HashIDSalt)
}
