package src

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	len1 := getLen(l1)
	len2 := getLen(l2)

	if len1 < len2 {
		l1, l2 = l2, l1
		len1, len2 = len2, len1
	}

	cur1 := l1
	cur2 := l2
	for i := len1 - len2; i > 0; i-- {
		cur1 = cur1.Next
	}

	for cur1 != nil {
		cur1.Val += cur2.Val
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	l1 = reverse(l1)
	carryBits(l1)
	l1 = reverse(l1)
	return l1
}

func getLen(l *ListNode) int {
	if l == nil {
		return 0
	}
	length := 1
	for l.Next != nil {
		l = l.Next
		length++
	}
	return length
}

func reverse(l *ListNode) *ListNode {
	var last *ListNode
	for l != nil {
		next := l.Next
		l.Next = last
		last = l
		l = next
	}
	return last
}

func carryBits(l *ListNode) {
	cur := l
	carry := false
	for cur != nil {
		if carry {
			carry = false
			cur.Val += 1
		}

		if cur.Val >= 10 {
			cur.Val -= 10
			carry = true
		}

		cur = cur.Next
	}
}

func TestQ2(t *testing.T) {
	l1 := &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{Val: 3},
		},
	}

	l2 := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{Val: 4},
		},
	}

	res := addTwoNumbers(l1, l2)
	fmt.Printf("%v", res)
}
