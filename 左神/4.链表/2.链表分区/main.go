package main

import "fmt"

/**
将单向链表按某值划分成左边小、中间相等、右边大的形式

【题目】给定一个单链表的头节点head，节点的值类型是整型，再给定一个整
数pivot 。实现一个调整链表的函数，将链表调整为左部分都是值小于pivot的
节点，中间部分都是值等于pivot的节点，右部分都是值大于pivot的节点。

【进阶】在实现原问题功能的基础上增加如下的要求

【要求】调整后所有小于pivot的节点之间的相对顺序和调整前一样
【要求】调整后所有等于pivot的节点之间的相对顺序和调整前一样
【要求】调整后所有大于pivot的节点之间的相对顺序和调整前一样
【要求】时间复杂度请达到0(N) ，额外空间复杂度请达到0(1) 。


*/
func main() {

	n := &Node{
		Val: 0,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 6,
				Next: &Node{
					Val: 2,
					Next: &Node{
						Val: 1,
						Next: &Node{
							Val: 10,
						},
					},
				},
			},
		},
	}

	fmt.Println(partition(n, 2))

}

type Node struct {
	Val  int
	Next *Node
}

// 链表转数组
func partition(head *Node, pivot int) *Node {

	arr := make([]*Node, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}

	start := -1
	end := len(arr)
	for i := 0; i < len(arr) && i < end; {
		if arr[i].Val < pivot {
			start++
			arr[start], arr[i] = arr[i], arr[start]
			i++
		} else if arr[i].Val == pivot {
			i++
		} else {
			end--
			arr[end], arr[i] = arr[i], arr[end]
		}
		// fmt.Println(i, func(arr []*Node) (vals []int) {
		// 	for _, item := range arr {
		// 		vals = append(vals, item.Val)
		// 	}
		// 	return
		// }(arr), start, end)
	}

	rn := arr[0]
	nn := rn
	for i := 1; i < len(arr); i++ {
		nn.Next = arr[i]
		nn = nn.Next
	}

	return rn

}

// 6指针
func partitionA(head *Node, pivot int) *Node {
	return nil
}
