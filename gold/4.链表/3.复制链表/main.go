package main

/**
复制含有随机指针节点的链表
【题目】一种特殊的单链表节点类描述如下

class Node {
	int value;
	Node next;
	Node rand;
	Node (int val) {
		value = val;
	}
}
rand指针是单链表节点结构中新增的指针，rand可能指向链表中的任意一个节
点，也可能指向nul1。给定一个由Node节点类型组成的无环单链表的头节点
head，请实现一个函数完成这个链表的复制，并返回复制的新链表的头节点。

【要求】时间复杂度0 (N) ，额外空间复杂度0 (1)

*/
func main() {

}

type Node struct {
	Val  int
	Next *Node
	Rand *Node
}

func copyByMap(head *Node) *Node {

	m := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		m[cur] = &Node{
			Val: cur.Val,
		}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Rand = m[cur.Rand]
		cur = cur.Next
	}

	return m[head]
}
