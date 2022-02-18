package main

import "fmt"

/**
题目: 输入一个链表的头节点，从尾到头反过来打印出每个节点的值
*/
func main() {

	node := &Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 3,
			},
		},
	}

	// nn := printRevLinkedA(node)
	nn := printRevLinkedB(node)

	for nn != nil {
		fmt.Println(nn.Val)
		nn = nn.Next
	}

}

type Node struct {
	Val  int
	Next *Node
}

// 头插法
func printRevLinkedA(node *Node) *Node {

	var nn *Node
	for node != nil {
		nn = &Node{
			Val:  node.Val,
			Next: nn,
		}
		node = node.Next
	}
	return nn
}

// 临时变量
func printRevLinkedB(node *Node) *Node {
	var prev *Node
	cur := node
	for cur != nil {
		temp := cur.Next

		cur.Next = prev
		prev = cur

		cur = temp
	}
	return prev
}
