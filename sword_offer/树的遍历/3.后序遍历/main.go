package main

import "fmt"

/**
后序遍历，左右根
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
	fmt.Println(preorderC(tree))

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
	vals = append(vals, preorderA(node.Left)...)
	// 中
	vals = append(vals, preorderA(node.Right)...)
	// 后
	vals = append(vals, node.Val)

	return

}

// 迭代实现
func preorderB(node *Node) (vals []int) {
	stack := make([]*Node, 0)
	cur := node
	var prev *Node

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.Right == nil || cur.Right == prev {
			vals = append(vals, cur.Val)
			prev = cur
			cur = nil
		} else {
			stack = append(stack, cur)
			cur = cur.Right
		}

	}

	return
}

// 迭代实现
func preorderC(node *Node) (vals []int) {
	stack := make([]*Node, 0)
	reverse := make([]*Node, 0)
	cur := node
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			reverse = append(reverse, cur)
			stack = append(stack, cur)
			cur = cur.Right
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Left
	}

	for i := len(reverse) - 1; i >= 0; i-- {
		vals = append(vals, reverse[i].Val)
	}

	return
}
