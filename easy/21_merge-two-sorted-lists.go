package main

import (
	"strconv"
	"strings"
)

/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

*/

//func main() {
//	a1 := &ListNode{Val: 1}
//	a2 := &ListNode{Val: 7}
//	a3 := &ListNode{Val: 8}
//
//	a1.Next = a2
//	a2.Next = a3
//
//	b1 := &ListNode{Val: 1}
//	b2 := &ListNode{Val: 2}
//	b3 := &ListNode{Val: 3}
//	b4 := &ListNode{Val: 4}
//
//	b1.Next = b2
//	b2.Next = b3
//	b3.Next = b4
//
//	fmt.Println(a1)
//	fmt.Println(b1)
//
//	fmt.Println(mergeTwoLists(a1, b1))
//
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	str := strconv.Itoa(l.Val) + "->"
	next := l.Next
	for {
		str += strconv.Itoa(next.Val) + "->"
		next = next.Next
		if next == nil {
			str = strings.TrimSuffix(str, "->")
			return str
		}
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	next := &res
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			*next = l2
			l2 = l2.Next
		} else {
			*next = l1
			l1 = l1.Next
		}
		next = &(*next).Next
	}
	if l1 == nil {
		*next = l2
	}
	if l2 == nil {
		*next = l1
	}
	return res
}
