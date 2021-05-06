package main

import (
	"fmt"
	"time"
)

//bubbleSort O(n^2)
func bubbleSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func recursionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
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
	fmt.Println(bubbleSort([]int{10, 45, 10, 6, 60, 31, 87, 29, 96, 6, 92, 81, 99, 72, 91, 82, 68, 33, 20, 90, 91, 67, 36, 26, 2, 70, 58, 6, 67, 91}))
	fmt.Println(recursionSort([]int{10, 45, 10, 6, 60, 31, 87, 29, 96, 6, 92, 81, 99, 72, 91, 82, 68, 33, 20, 90, 91, 67, 36, 26, 2, 70, 58, 6, 67, 91}))
	fmt.Println("耗时：", time.Since(start))
}
