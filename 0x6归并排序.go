package main

import (
	"fmt"
	"time"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	return merge(mergeSort(arr[0:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	arr := make([]int, 0)
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			arr = append(arr, left[0])
			left = left[1:]
		} else {
			arr = append(arr, right[0])
			right = right[1:]
		}
		i++
	}
	for len(left) > 0 {
		arr = append(arr, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		arr = append(arr, right[0])
		right = right[1:]
	}
	return arr
}

func main() {
	start := time.Now()
	fmt.Println(mergeSort([]int{10, 45, 10, 6, 60, 31, 87, 29, 96, 6, 92, 81, 99, 72, 91, 82, 68, 33, 20, 90, 91, 67, 36, 26, 2, 70, 58, 6, 67, 91}))
	fmt.Println("计算耗时: ", time.Since(start))
}
