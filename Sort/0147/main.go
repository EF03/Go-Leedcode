package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ln3 := ListNode{Val: 3, Next: nil}
	ln1 := ListNode{Val: 1, Next: &ln3}
	ln2 := ListNode{Val: 2, Next: &ln1}
	ln4 := ListNode{Val: 4, Next: &ln2}
	head := ListNode{Val: 0, Next: &ln4}
	fmt.Printf("===================== 开始 ===================\n")
	//loop(&head)
	newHead := insertionSortList(&head)

	fmt.Printf("===================== 结束 ===================\n")
	loop(newHead)
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: nil} // 这里初始化不要直接指向 head，为了下面循环可以统一处理
	cur, pre := head, newHead
	for cur != nil {
		// 记录下一个 loop用
		next := cur.Next
		// loop 找到比pre 比cur.Val 小一号的pre
		for pre.Next != nil && pre.Next.Val < cur.Val {
			fmt.Printf("###  pre = %v, next = %v\n  ", pre, pre.Next)
			pre = pre.Next
		}
		// 接上cur 替换pre原有的位置
		cur.Next = pre.Next
		pre.Next = cur
		fmt.Printf("end ===  pre = %v, cur = %v, next = %v\n  ", pre, cur, next)
		pre = newHead // 归位，重头开始
		cur = next
	}
	return newHead.Next
}

func loop(head *ListNode) {
	if head == nil {
		fmt.Printf("头是空的\n")
		return
	}
	cur := head
	// 循环
	for cur != nil {
		fmt.Printf("type = %T ,cur = %v\n  ", cur, cur)
		next := cur.Next
		cur = next
	}
}
