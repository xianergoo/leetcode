package two_pointers

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{1, 2, -2, -1}, [][]int{}},
	}
	for _, tt := range tests {
		got := threeSum(tt.nums)
		if !equalGroup(got, tt.want) {
			t.Errorf("threeSum(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

func equalGroup(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	sign := func(g [][]int) map[string]bool {
		m := make(map[string]bool)
		for _, v := range g {
			sort.Ints(v)
			m[fmt.Sprint(v)] = true
		}
		return m
	}
	sa, sb := sign(a), sign(b)
	return reflect.DeepEqual(sa, sb)
}
