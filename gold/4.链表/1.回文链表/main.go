package main

import "fmt"

/**
判断一个链表是否为回文结构

【题目】给定一个单链表的头节点head，请判断该链表是否为回文结构。
【例子】1->2->1，返回true; 1->2->2->1，返回true;i 15->6->15，返回true;
1->2->3，返回false。

【例子】如果链表长度为N，时间复杂度达到0(N) ，额外空间复杂度达到0(1) 。

*/
func main() {

	n := &Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 6,
				Next: &Node{
					Val: 2,
					Next: &Node{
						Val: 1,
					},
				},
			},
		},
	}

	// fmt.Println(palindromeA(n))
	fmt.Println(palindromeB(n))

}

type Node struct {
	Val  int
	Next *Node
}

// 利用栈
func palindromeA(node *Node) bool {

	stack := make([]*Node, 0)
	cur := node
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}

	flag := true
	cur = node
	for cur != nil && len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Val != pop.Val {
			flag = false
			break
		}
		cur = cur.Next
	}

	return flag

}

// 快慢指针+利用栈
func palindromeB(node *Node) bool {

	slow := node
	fast := node
	for fast != nil {
		if fast.Next == nil {
			// 奇数个中点
			break
		} else if fast.Next.Next == nil {
			// 偶数个中点
			slow = slow.Next
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	stack := make([]*Node, 0)
	for slow != nil {
		stack = append(stack, slow)
		slow = slow.Next
	}

	flag := true
	cur := node
	for cur != nil && len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pop.Val != cur.Val {
			flag = false
			break
		}
		cur = cur.Next
	}

	return flag
}

// 快慢指针+反转（恢复）
func palindromeC(node *Node) bool {
	// todo
	return true
}
