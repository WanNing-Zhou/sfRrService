package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXkiOiJhcHAiLCJleHAiOjE3MTEzMDI4NjMsIm5iZiI6MTcxMTI1ODY2MywiaWF0IjoxNzExMjU5NjYzLCJqdGkiOiI1MDU5MjMzMjg1ODY4MDk5MTAifQ.yzIhEHWu1-PjywVI_XPsYI72Rtwh5yLtVRztbJyGKWc"
	bearerToken, rest := splitToken(input)
	fmt.Println("Bearer Token:", bearerToken)
	fmt.Println("Rest:", rest)
}

func splitToken(input string) (string, string) {
	bearerPos := strings.Index(input, ".")
	return input[:bearerPos], input[bearerPos+1:]
}
