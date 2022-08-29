package main

import (
	"fmt"
	"math"
)

/**
如何判断一颗二叉树是否是搜索二又树?
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

	fmt.Println(isSearchTwoTree(tree))

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 迭代实现
func isSearchTwoTree(node *Node) bool {

	stack := make([]*Node, 0)

	compare := math.MinInt
	flag := true

	cur := node
	for cur != nil || len(stack) > 0 {

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if compare > cur.Val {
			flag = false
			break
		}
		compare = cur.Val

		cur = cur.Right

	}

	return flag
}
