package main

/**
题目: 给定一棵二又树和其中的一个节点，如何找出中序遍历的下一个节点?
树中的节点除了有两个分别指向在、右子节点的指针，还有一个指向父节点的指针。
*/
func main() {

}

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func NextNode(node *TreeNode) *TreeNode {

	if node == nil {
		return nil
	}

	// 有右节点
	if current := node.Right; current != nil {
		for current != nil {
			current = current.Left
		}
		return current
	}

	// 没有右节点
	current := node
	for current != nil {
		if current.Parent == nil {
			current = nil
			break
		} else if current.Parent.Left == current {
			break
		}
		current = current.Parent
	}

	return current
}
