package hash_table

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1}, 0, 0},
		{[]int{-1, -1, 1}, 0, 1},
		{[]int{1, 2, 1, 2, 1}, 3, 4},
	}
	for _, tt := range tests {
		got := subarraySum(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("subarraySum(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}
