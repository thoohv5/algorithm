package main

import "fmt"

/**
层级遍历，广度优先
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

	fmt.Println(preorder(tree))

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 迭代实现
func preorder(node *Node) (vals []int) {
	stack := make([]*Node, 0)

	cur := node
	stack = append(stack, cur)
	for i := 0; i < len(stack); i++ {
		cur = stack[i]
		vals = append(vals, cur.Val)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	return
}
