package midding

/**
* @Author: ZwZ
* @Time : 2020/9/28 15:24
* @File : LengthOfLongestSubstring
* @Description:
给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。
 * <p>
 * 示例1:
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * 示例 2:
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * 示例 3:
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
 * <p>
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
*/
func lengthOfLongestSubstring(s string) int {
	//哈希集合 key为元素 value为此字符出现的次数
	m := make(map[byte]int)
	length := len(s)
	right, res := -1, 0
	for i := 0; i < length; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right+1 < length && m[s[right+1]] == 0 {
			//移动右指针
			m[s[right+1]]++
			right++
		}
		res = max(res, right-i+1)
	}
	return res
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
