package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	m := [3]string{"()", "{}", "[]"}
	count := len(s)
	if count%2 != 0 {
		return false
	}
	i := 0
	for i < count/2 {
		for _, v := range m {
			s = strings.ReplaceAll(s, v, "")
		}
		i++
	}
	return s == ""
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
