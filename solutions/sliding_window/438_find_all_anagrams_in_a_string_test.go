package sliding_window

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
		{"a", "ab", []int{}},
		{"baa", "aa", []int{1}},
	}

	for _, tt := range tests {
		got := findAnagrams(tt.s, tt.p)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findAnagrams(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
		}
	}
}
