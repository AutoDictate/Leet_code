package main

import "fmt"

func AddTwoNumbersInit() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	result := AddTwoNumbers(l1, l2)
	for result != nil {
		fmt.Print(result.Val)
		if result.Next != nil {
			fmt.Print(" -> ")
		}
		result = result.Next
	}
	fmt.Println()
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1.Next != nil || l2.Next != nil {

		x, y := 0, 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry
		carry = sum / 10
		sum = sum % 10

		current.Next = &ListNode{
			Val: sum,
		}

		current = current.Next

		if carry > 0 {
			current.Next = &ListNode{
				Val: carry,
			}
		}

	}

	return dummy.Next

}