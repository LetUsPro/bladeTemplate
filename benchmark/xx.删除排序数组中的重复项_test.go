package benchmark

import "testing"

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mp := make(map[int]int, len(nums))
	target := nums[:0]
	for _, num := range nums {
		if _, ok := mp[num]; !ok {
			mp[num] = 1
			target = append(target, num)
		}
	}
	return len(target)
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicates([]int{0, 1, 2, 1, 3, 5, 2})
	}
}
