package hash_table

import (
	"fmt"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			[]string{""},
			[][]string{{""}},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		got := groupAnagrams(tt.strs)
		if !equalIgnoringOrder(got, tt.want) {
			t.Errorf("groupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
		}
	}
}

// 忽略组内和组间顺序比较两个 [][]string
func equalIgnoringOrder(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	// 签名：每组排序后拼成字符串作为 map key
	sign := func(groups [][]string) map[string]bool {
		m := make(map[string]bool)
		for _, g := range groups {
			sorted := make([]string, len(g))
			copy(sorted, g)
			sort.Strings(sorted)
			key := fmt.Sprint(sorted)
			m[key] = true
		}
		return m
	}
	sa, sb := sign(a), sign(b)
	if len(sa) != len(sb) {
		return false
	}
	for k := range sa {
		if !sb[k] {
			return false
		}
	}
	return true
}
