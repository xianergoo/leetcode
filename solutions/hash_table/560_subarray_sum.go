package hash_table

// 题号: 560
// 题目: 和为 K 的子数组 (Subarray Sum Equals K)
// 难度: Medium
// 链接: https://leetcode.cn/problems/subarray-sum-equals-k/
//
// 题目描述:
// 给你一个整数数组 nums 和一个整数 k，统计并返回该数组中和为 k 的连续子数组的个数。
//
// 示例:
//   输入: nums = [1,1,1], k = 2
//   输出: 2
//
//   输入: nums = [1,2,3], k = 3
//   输出: 2
//
// 思路:
// 前缀和 + 哈希表
//   preSum[j] - preSum[i] = k  →  preSum[i] = preSum[j] - k
//   遍历时维护 map[前缀和]出现次数，在 map 中查找 preSum - k 的出现次数累加
//
// 时间复杂度: O(n)
// 空间复杂度: O(n)

func subarraySum(nums []int, k int) int {
	pre, count := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i ++ {
		pre += nums[i]
		if _, ok := m[pre - k]; ok {
			count += m[pre - k]
		}
		m[pre] += 1
	}
	return count
}
