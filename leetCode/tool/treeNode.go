package tool

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func (this *TreeNode) Create(val int) {
	this.Val = val
}
