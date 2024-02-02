package util

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func RandStr(length int) string {
	res := make([]byte, (length+1)/2)
	rand.Read(res)

	return fmt.Sprintf("%x", res)
}

func StrWithDefault(source, defaultStr string) string {
	if source == "" {
		return defaultStr
	}

	return source
}

func Str2Bool(data string) bool {

	return strings.Compare(strings.ToUpper(data), "TRUE") == 0
}
