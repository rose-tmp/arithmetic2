package sort

/**
* @Author: ZwZ
* @Time : 2020/11/17 15:07
* @File : bubbleSort
* @Description:
 */
/**
原始冒泡排序
*/
func BubbleSort1(arr []int) {
	if len(arr) == 0 {
		return
	}
	for end := len(arr) - 1; end > 0; end-- {
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				Swap(arr, i, i+1)
			}
		}
	}
}

/**
改进后的冒泡排序
*/
func BubbleSort2(arr []int) {
	if len(arr) == 0 {
		return
	}
	for end := len(arr) - 1; end > 0; end-- {
		flag := true
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				Swap(arr, i, i+1)
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
func Swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
