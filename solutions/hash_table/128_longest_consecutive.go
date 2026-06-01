package hash_table

// 题号: 128
// 题目: 最长连续序列 (Longest Consecutive Sequence)
// 难度: Medium
// 链接: https://leetcode.cn/problems/longest-consecutive-sequence/
//
// 题目描述:
// 给定一个未排序的整数数组 nums，找出数字连续的最长序列的长度。
// 要求时间复杂度 O(n)。
//
// 示例:
//   输入: nums = [100,4,200,1,3,2]
//   输出: 4
//   解释: 最长连续序列是 [1,2,3,4]
//
// 思路:
// 1. 全部数放入 map[int]bool 当集合，O(1) 判断存在性
// 2. 只从「序列起点」（num-1 不在集合中）开始向后数
// 3. 更新全局最大长度
//
// 时间复杂度: O(n)   （每个数最多被访问2次：一次判断起点，一次往后数）
// 空间复杂度: O(n)

func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		m[v] = true
	}
	longest := 0
	for v := range m {
		if !m[v-1] {
			cur := v
			for m[cur] {
				cur++
			}
			if cur-v > longest {
				longest = cur - v
			}
		}
	}
	return longest
}
