package midding

import "arithmetic/src/leetCode/tool"

/**
* @Author: ZwZ
* @Time : 2020/11/24 14:46
* @File : countNodes
* @Description:222.完全二叉树的节点个数
* 要求时间复杂度低于O(N)【前中后遍历的时间复杂度都是O(N)】
*
* 高度为h的满二叉树的节点个数：2^h - 1
*
* 思路：
* 求出该树的深度h，(一直向左走，走到底)
* 从根节点的右子树出发向左走，走到底的高度为h2
* 如果h2 = h,那么根节点的左子树一定是满的，直接使用满二叉树节点个数的公式求解，然后再求右子树的高度(其就是一个子问题，所以可以用递归解决)
* 如果h2 < h,那么根节点的右子树一定是满的，直接使用满二叉树节点个数公式求解，然后递归求左子树的节点个数
* 时间复杂度：这种办法，每一层都只遍历一个节点，一共有logN层
* 而每次遍历某个节点的时候都会沿着左节点走到底，所以整体的时间复杂度为：O((logN)^2)
 */
func countNodes(root *tool.TreeNode) int {
	if root == nil {
		return 0
	}
	return count(root)
}
func count(root *tool.TreeNode) int {
	if root == nil {
		return 0
	}
	high := getHigh(root)
	highR := getHigh(root.Right) + 1
	//此时根节点的左子树为满二叉树
	if highR == high {
		return (1 << (high - 1)) + count(root.Right)
	} else {
		return (1 << (highR - 1)) + count(root.Left)
	}
}
func getHigh(root *tool.TreeNode) int {
	high := 0
	for root != nil {
		high++
		root = root.Left
	}
	return high
}
