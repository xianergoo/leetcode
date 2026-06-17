package sliding_window

// 题号: 3
// 题目: 无重复字符的最长子串 (Longest Substring Without Repeating Characters)
// 难度: Medium
// 链接: https://leetcode.cn/problems/longest-substring-without-repeating-characters/
//
// 题目描述:
// 给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。
//
// 思路:
// 使用滑动窗口维护一个始终不含重复字符的区间 [left..right]：
// 1. 右指针向右扩张，将新字符加入窗口
// 2. 若新字符导致重复，则不断移动左指针，缩小窗口并删除左侧字符
// 3. 直到窗口重新恢复为无重复字符，再用当前窗口长度更新答案
//
// 时间复杂度: O(n)
// 空间复杂度: O(min(n, |Sigma|))

// ==================== 解法 ====================

func lengthOfLongestSubstring(s string) int {
	windows := make(map[byte]int)
	left, right := 0, 0
	ans := 0
	for right < len(s) {
		ch := s[right]
		windows[ch]++
		for windows[ch] > 1 {
			windows[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}
