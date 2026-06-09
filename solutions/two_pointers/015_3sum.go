package two_pointers

import "sort"

// 题号: 15
// 题目: 三数之和 (3Sum)
// 难度: Medium
// 链接: https://leetcode.cn/problems/3sum/
//
// 思路:
// 排序 + 双指针。固定一个数，剩余范围用双指针找两数和。
//
// 时间复杂度: O(n²)
// 空间复杂度: O(log n)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
