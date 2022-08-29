package main

import (
	"fmt"
)

/**
如何判断一颗二叉树是否是完全二又树?
*/
func main() {

	tree := &Node{
		Val: 5,
		Left: &Node{
			Val: 3,
			Left: &Node{
				Val: 2,
			},
			Right: &Node{
				Val: 4,
			},
		},
		Right: &Node{
			Val: 7,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 9,
			},
		},
	}

	fmt.Println(isCompletelyTwoTree(tree))

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 迭代实现
func isCompletelyTwoTree(node *Node) bool {

	stack := make([]*Node, 0)

	cur := node
	stack = append(stack, cur)
	index := 0
	flag := true
	isLeaves := false
	for len(stack) > 0 && index < len(stack) {
		cur = stack[index]

		if isLeaves && !(cur.Left == nil && cur.Right == nil) || cur.Left == nil && cur.Right != nil {
			flag = false
			break
		} else if cur.Left != nil && cur.Right == nil {
			isLeaves = true
		}

		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	return flag
}
