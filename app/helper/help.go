package helper

import (
	"encoding/base64"
	"log"
)

func Encode64(s []byte) string {
	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	return s64
}
func Decode64(s string) string {
	bs, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalln(err)
	}
	return string(bs)
}
