package sliding_window

// 题号: 438
// 题目: 找到字符串中所有字母异位词 (Find All Anagrams in a String)
// 难度: Medium
// 链接: https://leetcode.cn/problems/find-all-anagrams-in-a-string/
//
// 题目描述:
// 给定两个字符串 s 和 p，找到 s 中所有 p 的异位词子串，返回这些子串的起始索引。
//
// 思路:
// 使用固定长度滑动窗口维护当前子串的字符频次：
// 1. 先统计 p 的字符频次 need，并初始化 s 的第一个长度为 len(p) 的窗口频次
// 2. 若当前窗口频次与 need 相同，则当前窗口起点是一个答案
// 3. 窗口每次右移一格：移出左边字符，移入右边字符，再比较两张频次表
//
// 时间复杂度: O(n)
// 空间复杂度: O(1)

// ==================== 解法 ====================

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	need := [26]int{}
	for i := 0; i < len(p); i++ {
		need[p[i]-'a']++
	}

	ans := make([]int, 0)
	windowsLen := len(p)
	cur := [26]int{}
	for i := 0; i < windowsLen; i++ {
		cur[s[i]-'a']++
	}
	if cur == need {
		ans = append(ans, 0)
	}
	left, right := 0, windowsLen
	for right < len(s) {
		cur[s[left]-'a']--
		cur[s[right]-'a']++
		if cur == need {
			ans = append(ans, left+1)
		}
		right++
		left++

	}
	return ans
}
