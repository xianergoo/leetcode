package hash_table

import "sort"

// 题号: 49
// 题目: 字母异位词分组 (Group Anagrams)
// 难度: Medium
// 链接: https://leetcode.cn/problems/group-anagrams/
//
// 题目描述:
// 给你一个字符串数组，请你将字母异位词组合在一起。可以按任意顺序返回结果列表。
// 字母异位词是字母相同，但排列不同的字符串。
//
// 示例:
//   输入: strs = ["eat","tea","tan","ate","nat","bat"]
//   输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 思路:
// 排序法 — 每个字符串按字符排序后作为 map 的 key，同组字符串聚到同一个 value 列表。
//
// 时间复杂度: O(n * k log k)  （n=字符串个数, k=字符串平均长度）
// 空间复杂度: O(n * k)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		key := sortString(s)
		m[key] = append(m[key], s)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func sortString(s string) string {
	r := []byte(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
