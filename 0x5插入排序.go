package main

import (
	"fmt"
	"time"
)

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i
		for j > 0 && tmp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		if i != j {
			arr[j] = tmp
		}
	}
	return arr
}

func main() {
	start := time.Now()
	fmt.Println(insertSort([]int{10, 45, 10, 6, 60, 31, 87, 29, 96, 6, 92, 81, 99, 72, 91, 82, 68, 33, 20, 90, 91, 67, 36, 26, 2, 70, 58, 6, 67, 91}))
	fmt.Println("计算耗时: ", time.Since(start))
}
