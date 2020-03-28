package main

/*
83. 删除排序链表中的重复元素
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:
输入: 1->1->2
输出: 1->2

示例 2:
输入: 1->1->2->3->3
输出: 1->2->3
*/

//func main() {
//	a1 := &ListNode{Val: 1}
//	a2 := &ListNode{Val: 1}
//	a3 := &ListNode{Val: 1}
//
//	a1.Next = a2
//	a2.Next = a3
//
//	b1 := &ListNode{Val: 1}
//	b2 := &ListNode{Val: 2}
//	b3 := &ListNode{Val: 4}
//	b4 := &ListNode{Val: 4}
//
//	b1.Next = b2
//	b2.Next = b3
//	b3.Next = b4
//
//	fmt.Println(a1)
//	fmt.Println(b1)
//
//	fmt.Println(deleteDuplicates2(a1))
//	fmt.Println(deleteDuplicates2(b1))
//}

// 没有审题，忘记是排序链表
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}

	next := head
	m := make(map[int]struct{})
	m[head.Val] = struct{}{}
	for next.Next != nil {
		if _, ok := m[next.Next.Val]; !ok {
			m[next.Next.Val] = struct{}{}
			next = next.Next
			continue
		}
		next.Next = next.Next.Next
	}
	return head
}

// 操作next指针就行
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	pre := head
	next := head.Next
	for next != nil {
		if pre.Val == next.Val {
			pre.Next = next.Next
			next = pre.Next
			continue
		}
		pre = next
		next = pre.Next
	}
	return head
}
