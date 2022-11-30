package helper

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5V 生存MD5字符串
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
