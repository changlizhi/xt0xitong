package ml0gongjus

import (
	"crypto/md5"
	"encoding/hex"
)

// 生成32位MD5
func Md5JiaMi(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
