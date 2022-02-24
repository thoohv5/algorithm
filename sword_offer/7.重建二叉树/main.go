package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	recursive(preorder, inorder, 0, len(preorder), 0, len(inorder), root)
	return root
}

func recursive(preorder, inorder []int, s, e, x, y int, tree *TreeNode) {
	if e == s {
		return
	}

	nv := preorder[s]

	index := -1
	for i := x; i < y; i++ {
		if inorder[i] == nv {
			index = i
			break
		}
	}
	tree.Val = inorder[index]

	if s+1 < s+1+index-x {
		tree.Left = &TreeNode{}
		recursive(preorder, inorder, s+1, s+1+index-x, x, index, tree.Left)
	}

	if s+1+index-x < e {
		tree.Right = &TreeNode{}
		recursive(preorder, inorder, s+1+index-x, e, index+1, y, tree.Right)
	}
}
