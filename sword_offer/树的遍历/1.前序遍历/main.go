package main

import "fmt"

/**
前序遍历，根左右
*/
func main() {

	tree := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 3,
			},
			Right: &Node{
				Val: 4,
			},
		},
		Right: &Node{
			Val: 5,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	fmt.Println(preorderA(tree))
	fmt.Println("---------")
	fmt.Println(preorderB(tree))

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 递归实现
func preorderA(node *Node) (vals []int) {

	if node == nil {
		return
	}

	// 前
	vals = append(vals, node.Val)
	vals = append(vals, preorderA(node.Left)...)
	// 中
	vals = append(vals, preorderA(node.Right)...)
	// 后

	return

}

// 迭代实现
func preorderB(node *Node) (vals []int) {
	stack := make([]*Node, 0)
	cur := node
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			vals = append(vals, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
	return
}
