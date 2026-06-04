package two_pointers

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
	}
	for _, tt := range tests {
		got := make([]int, len(tt.nums))
		copy(got, tt.nums)
		moveZeroes(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("moveZeroes(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
