package midding

import "arithmetic/src/leetCode/tool"

/**
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明:叶子节点是指没有子节点的节点。
示例:
给定如下二叉树，以及目标和sum = 22，
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
[
   [5,4,11,2],
   [5,8,4,5]
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-ii
*/
func pathSum(root *tool.TreeNode, sum int) (ans [][]int) {
	var path []int
	var dfs func(*tool.TreeNode, int)
	dfs = func(root *tool.TreeNode, left int) {
		if root == nil {
			return
		}
		left -= root.Val
		path = append(path, root.Val)
		defer func() { path = path[:len(path)-1] }()
		if root.Left == nil && root.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(root.Left, left)
		dfs(root.Right, left)
	}
	dfs(root, sum)
	return ans
	// return  函数返回值有变量名 即ans 所以不显示的指定返回值为ans 直接一个return也是可以的
}
