package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 加密
func MD5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
