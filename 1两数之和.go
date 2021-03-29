package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	m := make(map[int]int)

	for k, v := range nums {
		if n, o := m[target-v]; o {
			result = append(result, n, k)
			goto Loop
		}
		m[v] = k
	}
Loop:
	return result
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
