package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) (ret *ListNode) {
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = ret
		ret = cur
		cur = temp
	}
	return
}
