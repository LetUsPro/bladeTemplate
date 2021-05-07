package main

import (
	"fmt"
	"time"
)

func recursionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	less := make([]int, 0)
	great := make([]int, 0)
	equal := make([]int, 0)
	for _, v := range arr {
		if v < pivot {
			less = append(less, v)
		}
		if v > pivot {
			great = append(great, v)
		}
		if v == pivot {
			equal = append(equal, v)
		}
	}
	return append(append(recursionSort(less), equal...), recursionSort(great)...)
}

func main() {
	start := time.Now()
	fmt.Println(recursionSort([]int{10, 45, 10, 6, 60, 31, 87, 29, 96, 6, 92, 81, 99, 72, 91, 82, 68, 33, 20, 90, 91, 67, 36, 26, 2, 70, 58, 6, 67, 91}))
	fmt.Println("耗时：", time.Since(start))
}
