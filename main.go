package main

import (
	"arithmetic/leetCode/easy"
	"arithmetic/leetCode/greedy"
	"fmt"
)

/**
* @Author: ZwZ
* @Time : 2020/11/17 15:31
* @File : main
* @Description:
 */
func main() {
	/*var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.BubbleSort1(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}*/
	/*var arr = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	greedy.FindMinArrowShots(arr)*/
	str := "aaabbbccc"
	easy.SortString(str)
	fmt.Println(string(97))
	var mapD map[int]int
	mapD = make(map[int]int, 10)
	for i := 0; i < 12; i++ {
		mapD[i] = i
	}
	var arr = []int{1, 2, 3, 3, 4, 5}
	fmt.Println(greedy.IsPossible(arr))

}
