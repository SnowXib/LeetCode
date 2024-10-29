package medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reverse_slice(currentSlice []int) []int {
	returneSlice := []int{}
	for index := range currentSlice {
		returneSlice = append(returneSlice, currentSlice[len(currentSlice)-1-index])
	}
	return returneSlice
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1_slice := []int{}
	l2_slice := []int{}
	var diff_info bool = true

	for l1 != nil || l2 != nil {
		if l1 != nil {
			l1_slice = append(l1_slice, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			l2_slice = append(l2_slice, l2.Val)
			l2 = l2.Next
		}
	}

	l1_slice = reverse_slice(l1_slice)
	l2_slice = reverse_slice(l2_slice)

	diff := len(l1_slice) - len(l2_slice)
	if diff < 0 {
		diff_info = false
	}

	append_slice := make([]int, abs(diff), abs(diff))

	if diff_info {
		l2_slice = append(append_slice, l2_slice...)
	} else {
		l1_slice = append(append_slice, l1_slice...)
	}

	sum_sllice := make([]int, len(l2_slice), len(l2_slice))

	for index := range l1_slice {
		in_index := len(l1_slice) - index - 1
		sum_sllice[in_index] = l1_slice[in_index] + l2_slice[in_index] + sum_sllice[in_index]
		for {
			if sum_sllice[in_index] > 9 {
				sum_sllice[in_index] = sum_sllice[in_index] - 10
				if in_index == 0 {
					var digit = []int{1}
					sum_sllice = append(digit, sum_sllice...)
				} else {
					fmt.Println(in_index)
					sum_sllice[in_index-1]++
				}
			} else {
				break
			}
			in_index--
			if in_index < 0 {
				break
			}
		}
	}

	fmt.Println(sum_sllice)

	l_sum := createList(reverse_slice(sum_sllice))

	return l_sum
}

func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}
