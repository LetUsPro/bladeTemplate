package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s1 := strings.ReplaceAll(s, " ", "")
	if s1 == "" {
		return 0
	}
	strArr := strings.Split(strings.Trim(s, " "), " ")
	fmt.Println(strArr)
	return len(strArr[len(strArr)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("hello "))
}
