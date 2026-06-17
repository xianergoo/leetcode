package sliding_window

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"au", 2},
	}

	for _, tt := range tests {
		got := lengthOfLongestSubstring(tt.s)
		if got != tt.want {
			t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
