package hash_table

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{6, 5, 5}, 5},
	}
	for _, tt := range tests {
		got := majorityElement(tt.nums)
		if got != tt.want {
			t.Errorf("majorityElement(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
