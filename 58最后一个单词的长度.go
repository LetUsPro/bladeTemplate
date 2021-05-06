package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	sli := strings.Split(s, " ")
	return len(sli[len(sli)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("hello "))
}
