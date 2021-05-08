package main

import (
	"fmt"
	"time"
)

func countSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	max := getMinMaxValue(arr)
	length := max + 1
	bucket := make([]int, length)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	index := 0
	for j := 0; j < length; j++ {
		for bucket[j] > 0 {
			arr[index] = j
			index++
			bucket[j]--
		}
	}
	return arr
}

func getMinMaxValue(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	start := time.Now()
	fmt.Println(countSort([]int{10, 45, 10, 6, 60, 31, 87, 29, 96, 6, 92, 81, 99, 72, 91, 82, 68, 33, 20, 90, 91, 67, 36, 26, 2, 70, 58, 6, 67, 91}))
	fmt.Println("计算耗时: ", time.Since(start))
}
