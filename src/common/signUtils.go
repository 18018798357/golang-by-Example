package common

import (
	"crypto/md5"
	"encoding/hex"
)

const secret string = "44d6d569341947ec947c711a18574de5"

// md5 加密
func CryptographyMd5 (parm string) string {
	md5 := md5.New()
	md5.Write([]byte(parm))
	md5Data := md5.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}
