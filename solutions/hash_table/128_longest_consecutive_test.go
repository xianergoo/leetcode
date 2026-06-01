package hash_table

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1, 1}, 1},
	}
	for _, tt := range tests {
		got := longestConsecutive(tt.nums)
		if got != tt.want {
			t.Errorf("longestConsecutive(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
