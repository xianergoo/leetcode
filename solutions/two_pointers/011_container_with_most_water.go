package two_pointers

// 题号: 11
// 题目: 盛最多水的容器 (Container With Most Water)
// 难度: Medium
// 链接: https://leetcode.cn/problems/container-with-most-water/
//
// 题目描述:
// 给定一个长度为 n 的整数数组 height。有 n 条垂线，找出两条线使得它们与 x 轴共同构成的容器
// 可以容纳最多的水。返回容器可以储存的最大水量。
//
// 思路:
// 双指针从两端向中间移动，每次移动较矮的那一边，计算并更新最大面积。
//
// 时间复杂度: O(n)
// 空间复杂度: O(1)

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		area = max(area, min(height[left], height[right])*(right-left))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return area
}
