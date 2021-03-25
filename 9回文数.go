package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || x > 0 && x%10 == 0 {
		return false
	}
	last := 0
	for x > last {
		last = last*10 + x%10
		x = x / 10
	}
	return x == last || x == last/10
}

func main() {
	fmt.Println(isPalindrome(121))
}
