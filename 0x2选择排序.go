package main

import (
	"fmt"
	"time"
)

func selectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	arrSlice := arr[:]
	ret := make([]int, 0)
	for _ = range arrSlice {
		smallIndex := findSmall(arrSlice)
		ret = append(ret, arr[smallIndex])
		arrSlice = append(arrSlice[0:smallIndex], arrSlice[smallIndex+1:]...)
	}
	return ret
}

func findSmall(arr []int) int {
	small := arr[0]
	smallIndex := 0
	for i, v := range arr {
		if v < small {
			smallIndex = i
		}
	}
	return smallIndex
}

func main() {
	start := time.Now()
	fmt.Println(selectSort([]int{1, 5, 8, 2, 100, 77, 35}))
	fmt.Printf("耗时计算: %s\n", time.Since(start))
}
