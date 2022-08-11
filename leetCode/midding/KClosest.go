package midding

import "container/heap"

/**
* @Author: ZwZ
* @Time : 2020/11/10 20:04
* @File : KClosestt
* @Description:973. 最接近原点的 K 个点
* 我们有一个由平面上的点组成的列表 points。需要-从中找出 K 个距离原点 (0, 0) 最近的点。
* （这里，平面上两点之间的距离是欧几里德距离。）
* 你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。
* <p>
* 示例 1：
* 输入：points = [[1,3],[-2,2]], K = 1
* 输出：[[-2,2]]
* 解释：
* (1, 3) 和原点之间的距离为 sqrt(10)，
* (-2, 2) 和原点之间的距离为 sqrt(8)，
* 由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
* 我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。
* <p>
* 示例 2：
* 输入：points = [[3,3],[5,-1],[-2,4]], K = 2
* 输出：[[3,3],[-2,4]]
* （答案 [[-2,4],[3,3]] 也会被接受。）
* <p>
* 来源：力扣（LeetCode）
* 链接：https://leetcode-cn.com/problems/k-closest-points-to-origin
 */
type pair struct {
	distance int
	point    []int
}
type hp []pair

/*
* 只有实现了这些方法的数据类型才可以加入堆中
 */
//绑定len方法，返回长度
func (h hp) Len() int {
	return len(h)
}

//绑定Less方法 如果h[i] < h[j]生成小根堆   这里我们要创建大根堆
func (h hp) Less(i, j int) bool {
	return h[j].distance < h[i].distance
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

//插入新元素
func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

//从最后拿出一个元素并返回
func (h *hp) Pop() interface{} {
	a := *h
	x := a[len(a)-1]
	*h = a[0 : len(a)-1]
	return x
}

func kClosest(points [][]int, K int) (ans [][]int) {
	h := make(hp, K) //创建slice
	//将points中下标0--K-1的元素的距离先加入堆
	for i, p := range points[:K] {
		h[i] = pair{p[0]*p[0] + p[1]*p[1], p}
	}

	heap.Init(&h) //使用切片h中的元素 初始化一个大小为K的堆
	for _, p := range points[K:] {
		if dist := p[0]*p[0] + p[1]*p[1]; dist < h[0].distance {
			h[0] = pair{distance: dist, point: p}
			heap.Fix(&h, 0) //效率比pop后再push高
		}
	}
	for _, p := range h {
		ans = append(ans, p.point)
	}
	return
}
