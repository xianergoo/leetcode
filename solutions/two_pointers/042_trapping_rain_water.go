package two_pointers

// 题号: 42
// 题目: 接雨水 (Trapping Rain Water)
// 难度: Hard
// 链接: https://leetcode.cn/problems/trapping-rain-water/
//
// 题目描述:
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 思路:
// 双指针维护左右两侧见过的最高柱子：
// 1. leftMax 表示左指针左侧的最高柱子，rightMax 表示右指针右侧的最高柱子
// 2. 哪一侧当前柱子更低，就先结算哪一侧，因为这一侧的积水上限已经确定
// 3. 若当前柱子低于这一侧最高柱子，则可接雨水 = maxHeight - height[i]
//
// 时间复杂度: O(n)
// 空间复杂度: O(1)

func trap(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}
	return ans
}
