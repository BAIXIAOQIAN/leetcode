/*
给定两个非空链表，用来表示两个非负整数，数字存储的顺序是反的，例如：2->4->5表示 542,
将两个非空链表表示的非负整数相加之和输出到一个链表中，存储的形式同上
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1, l2 *ListNode
	l1 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	result := addTwoNumbers(l1, l2)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	carry, current := 0, result
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		sum := carry + x + y
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}
	return result.Next
}

//result 的值从哪里来的，是因为指针指向同一个地址？
