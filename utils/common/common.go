package common

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func Strftime(t time.Time) string{
	return t.String()[:19]
}


func Strptime(t string, format string) (time.Time, error){
	structT, err := time.Parse(format, t)
	return structT, err
}


// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
