package hash_table

// 题号: 1
// 题目: 两数之和 (Two Sum)
// 难度: Easy
// 链接: https://leetcode.cn/problems/two-sum/
//
// 题目描述:
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值的那两个整数，
// 并返回它们的数组下标。你可以假设每种输入只会对应一个答案，且同一元素不能使用两遍。
//
// 思路:
// 哈希表存储已遍历元素的值到索引的映射，遍历时查找 target-nums[i] 是否已存在。
//
// 时间复杂度: O(n)
// 空间复杂度: O(n)

// ==================== 解法 ====================

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}


