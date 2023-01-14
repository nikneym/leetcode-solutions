/*
 * https://leetcode.com/problems/add-two-numbers/
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ll := ListNode{}
	ptr := &ll
	carry := 0

	for l1 != nil || l2 != nil || carry == 1 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10

		ptr.Next = &ListNode{
			Val: sum % 10,
		}
		ptr = ptr.Next
	}

	return ll.Next
}

func linkedListFrom(nums []uint) *ListNode {
	res := ListNode{}
	ptr := &res
	for _, v := range nums {
		*ptr = ListNode{
			Val:  int(v),
			Next: &ListNode{},
		}
		ptr = ptr.Next
	}

	return &res
}

func main() {
	l1 := linkedListFrom([]uint{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 := linkedListFrom([]uint{5, 6, 4})

	res := addTwoNumbers(l1, l2)
	list := make([]int, 0)
	for node := res; node != nil; node = node.Next {
		list = append(list, node.Val)
	}

	fmt.Println(list)
}
