package easy

import "sort"

/**
* @Author: ZwZ
* @Time : 2020/11/22 15:43
* @File : isAnagram
* @Description:242. 有效的字母异位词
* 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
* <p>
* 示例1:
* 输入: s = "anagram", t = "nagaram"
* 输出: true
* <p>
* 示例 2:
* 输入: s = "rat", t = "car"
* 输出: false
* <p>
* 说明:
* 你可以假设字符串只包含小写字母。
* <p>
* 来源：力扣（LeetCode）
* 链接：https://leetcode-cn.com/problems/valid-anagram
 */
func isAnagram1(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	return string(s1) == string(s2)
}
func isAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}
