package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	t := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	r := 0
	for k, n := range t {
		if strings.Contains(s, k) {
			r += n
			s = strings.ReplaceAll(s, k, "")
		}
	}
	for i := 0; i < len(s); i++ {
		r += m[s[i]]
	}
	return r
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
