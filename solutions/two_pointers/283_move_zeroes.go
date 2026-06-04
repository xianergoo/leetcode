package two_pointers

// 题号: 283
// 题目: 移动零 (Move Zeroes)
// 难度: Easy
// 链接: https://leetcode.cn/problems/move-zeroes/
//
// 题目描述:
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 必须在原数组上操作，不能拷贝额外的数组。
//
// 示例:
//   输入: nums = [0,1,0,3,12]
//   输出: [1,3,12,0,0]
//
// 思路:
// 双指针：用 pos 指向下一个非零元素该放的位置，遍历一遍将非零元素前移，
// 最后将 pos 之后的位置全部补 0。
//
// 时间复杂度: O(n)
// 空间复杂度: O(1)

func moveZeroes(nums []int) {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}
	for i := pos; i < len(nums); i++ {
		nums[i] = 0
	}
	return
}
