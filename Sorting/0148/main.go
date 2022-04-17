package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

var step int

func loopNode(l *ListNode) {
	for l != nil {
		fmt.Printf("val = %d \n", l.Val)
		l = l.Next
	}
}

func main() {
	node3 := ListNode{Val: 3}
	node1 := ListNode{Val: 1, Next: &node3}
	node2 := ListNode{Val: 2, Next: &node1}
	node4 := ListNode{Val: 4, Next: &node2}
	fmt.Println("=== 初始化 ===")
	loopNode(&node4)
	newHead := sortList(&node4)
	fmt.Println("=== result ===")
	loopNode(newHead)
}

func sortList(head *ListNode) *ListNode {
	step++
	fmt.Printf("%d #########################\n", step)
	if head == nil || head.Next == nil {
		return head
	}
	var slow, fast, pre = head, head, head
	// 快指针每次走两步，慢指针每次走一步，当快指针到达链表末尾时，慢指针正好走到中间
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	//fmt.Println("=== pre ===")
	//loopNode(pre)
	// 把兩 head slow 切開
	pre.Next = nil
	return mergeTwoLists(sortList(head), sortList(slow))

}

//                 left          right
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	step++
	fmt.Printf("%d #########################\n", step)
	// 2-4 |
	fmt.Println("=== mergeTwoLists l1 ===")
	loopNode(l1)
	// 1-3 |
	fmt.Println("=== mergeTwoLists l2 ===")
	loopNode(l2)
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		fmt.Println("=== merge l1 ===")
		loopNode(l1)
		return l1
	}
	//  1 , 2-3-4
	l2.Next = mergeTwoLists(l1, l2.Next)
	fmt.Println("=== merge l2 ===")
	loopNode(l2)
	return l2
}
