package main

import (
	"strconv"
	"strings"
)

/*
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

//func main() {
//	a1 := &ListNode{Val: 9}
//	a2 := &ListNode{Val: 9}
//	a3 := &ListNode{Val: 9}
//
//	a1.Next = a2
//	a2.Next = a3
//
//	b1 := &ListNode{Val: 1}
//	//b2 := &ListNode{Val: 6}
//	//b3 := &ListNode{Val: 4}
//
//	//b1.Next = b2
//	//b2.Next = b3
//
//	fmt.Println(a1)
//	fmt.Println(b1)
//
//	fmt.Println(addTwoNumbers(a1, b1))
//
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	str := ""
	for {
		str += strconv.Itoa(l.Val) + "->"
		l = l.Next
		if l == nil {
			str = strings.TrimSuffix(str, "->")
			return str
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	data := &ListNode{}
	next := data
	for {
		if l1 != nil {
			next.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			next.Val += l2.Val
			l2 = l2.Next
		}
		// 大于10，进位
		if (next.Val) >= 10 {
			next.Val -= 10
			next.Next = &ListNode{Val: 1}
		}

		if l1 == nil && l2 == nil {
			break
		}

		// 无进位直接创建
		if next.Next == nil {
			next.Next = &ListNode{}
		}
		next = next.Next
	}

	return data
}
