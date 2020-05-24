package Helper

import (
	"SWord/App"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

var writeStruct struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

func Write(code int, result interface{}, msg string) interface{} {
	write := writeStruct
	write.Code = code
	write.Result = result
	write.Msg = msg
	return write
}

func UsersPasswordEncrypt(password string) string {
	h1 := md5.New()
	h1.Write([]byte(App.SecretKey))
	secretKeyMd5 := hex.EncodeToString(h1.Sum(nil))

	var build strings.Builder
	build.WriteString(secretKeyMd5)
	build.WriteString(password)
	tmp := build.String()

	h2 := md5.New()
	h2.Write([]byte(tmp))
	rs := hex.EncodeToString(h2.Sum(nil))
	return rs
}
