package greedy

/**
* @Author: ZwZ
* @Time : 2020/12/4 15:19
* @File : isPossible
* @Description:659. 分割数组为连续子序列
* 给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个子序列，
 * 其中每个子序列都由连续整数组成且长度至少为 3。
 * 如果可以完成上述分割，则返回 true ；否则，返回 false 。
 *
 * 示例 1：
 * 输入: [1,2,3,3,4,5]
 * 输出: True
 * 解释:
 * 你可以分割出这样两个连续子序列 :
 * 1, 2, 3
 * 3, 4, 5
 *
 * 示例 2：
 * 输入: [1,2,3,3,4,4,5,5]
 * 输出: True
 * 解释:
 * 你可以分割出这样两个连续子序列 :
 * 1, 2, 3, 4, 5
 * 3, 4, 5
 *
 * 示例 3：
 * 输入: [1,2,3,4,4,5]
 * 输出: False
 *
 * 提示：
 * 输入的数组长度范围为 [1, 10000]
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences
*/
func IsPossible(nums []int) bool {
	if nums == nil {
		return false
	}
	var countMap, endMap map[int]int
	countMap = make(map[int]int)
	endMap = make(map[int]int)
	//初始化countMap
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]]++
	}
	//遍历nums的目的就是为其中的每个元素找到他们对应的一个子序列  即他们的归属
	for i := 0; i < len(nums); i++ {
		count := countMap[nums[i]]
		if count > 0 {
			//map中如果查找失败将返回value类型的0值
			preEnd := endMap[nums[i]-1]
			if preEnd > 0 {
				countMap[nums[i]]--
				endMap[nums[i]-1]--
				endMap[nums[i]]++
			} else {
				temp1 := countMap[nums[i]+1]
				temp2 := countMap[nums[i]+2]
				if temp1 > 0 && temp2 > 0 {
					countMap[nums[i]]--
					countMap[nums[i]+1]--
					countMap[nums[i]+2]--
					endMap[nums[i]+2]++
				} else {
					return false
				}
			}
		}
	}
	return true
}
