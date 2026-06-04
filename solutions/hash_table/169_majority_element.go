package hash_table

// 题号: 169
// 题目: 多数元素 (Majority Element)
// 难度: Easy
// 链接: https://leetcode.cn/problems/majority-element/
//
// 题目描述:
// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊n/2⌋ 的元素。
//
// 思路:
// 哈希表统计频次，遇 count > n/2 时返回
//
// 时间复杂度: O(n)
// 空间复杂度: O(n)

func majorityElement(nums []int) int {
	mp := make(map[int]int, len(nums))
	for _, v := range nums {
		mp[v]++
		if mp[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

func majorityElementBoyer(nums []int) int {
	candidate, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if candidate == v {
			count++
		} else {
			count--
		}
	}
	return candidate
}
