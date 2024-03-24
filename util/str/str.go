package str

import (
	"math/rand"
	"strings"
	"time"
)

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// SplitToken 拆分token
func SplitToken(input string) string {
	return strings.Replace(input, "Bearer ", "", -1)
	//return input[bearerPos+1:]
}
