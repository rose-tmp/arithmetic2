package easy

/**
* @Author: ZwZ
* @Time : 2020/11/25 19:53
* @File : sortString
* @Description:1370. 上升下降字符串
* 给你一个字符串s，请你根据下面的算法重新构造字符串：
 * 从 s中选出最小的字符，将它接在结果字符串的后面。
 * 从 s剩余字符中选出最小的字符，且该字符比上一个添加的字符大，将它接在结果字符串后面。
 * 重复步骤 2 ，直到你没法从 s中选择字符。
 * 从 s中选出最大的字符，将它 接在结果字符串的后面。
 * 从 s剩余字符中选出最大的字符，且该字符比上一个添加的字符小，将它 接在结果字符串后面。
 * 重复步骤 5，直到你没法从 s中选择字符。
 * 重复步骤 1 到 6 ，直到 s中所有字符都已经被选过。
 * 在任何一步中，如果最小或者最大字符不止一个，你可以选择其中任意一个，并将其添加到结果字符串。
 * <p>
 * 请你返回将s中字符重新排序后的 结果字符串 。
 * <p>
 * 示例 1：
 * 输入：s = "aaaabbbbcccc"
 * 输出："abccbaabccba"
 * 解释：第一轮的步骤 1，2，3 后，结果字符串为 result = "abc"
 * 第一轮的步骤 4，5，6 后，结果字符串为 result = "abccba"
 * 第一轮结束，现在 s = "aabbcc" ，我们再次回到步骤 1
 * 第二轮的步骤 1，2，3 后，结果字符串为 result = "abccbaabc"
 * 第二轮的步骤 4，5，6 后，结果字符串为 result = "abccbaabccba"
 * <p>
 * 示例 2：
 * 输入：s = "rat"
 * 输出："art"
 * 解释：单词 "rat" 在上述算法重排序以后变成 "art"
 * <p>
 * 示例 3：
 * 输入：s = "leetcode"
 * 输出："cdelotee"
 * <p>
 * 示例 4：
 * 输入：s = "ggggggg"
 * 输出："ggggggg"
 * <p>
 * 示例 5：
 * 输入：s = "spo"
 * 输出："ops"
 * <p>
 * 提示：
 * <p>
 * 1 <= s.length <= 500
 * s只包含小写英文字母。
 * <p>
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/increasing-decreasing-string
*/
func SortString(s string) string {
	if len(s) < 2 {
		return s
	}
	//桶
	var flag [26]int
	//Go没有专门的字符类型，如果要存储单个字符就要用byte去保存ASCII码
	//不能用make([]byte,len(s))因为这样初始化出的ans的长度和其cap一样 即len(s)  而每个位置上的数默认为0
	ans := make([]byte, 0, len(s))
	//初始化桶
	for _, ch := range s {
		flag[ch-'a']++
	}
	for len(ans) < len(s) {
		for i := 0; i < 26; i++ {
			if flag[i] > 0 {
				flag[i]--
				ans = append(ans, byte(i)+'a')
			}
		}
		for i := 25; i >= 0; i-- {
			if flag[i] > 0 {
				flag[i]--
				ans = append(ans, byte(i)+'a')
			}
		}
	}
	/**
	Go的字符串是由单个字节连接组成的，与传统的由字符组成的字符串不同
	使用string()可以将其转成字符串
	也可以使用string()转成某个字符  如string(97) 就是 a
	*/
	return string(ans)
}
