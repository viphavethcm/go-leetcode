package main

import (
	"fmt"
	"go-leetcode/algo/linked"
)

// Check string is number: 	_, err := strconv.Atoi(s)
// Convert string to number: 	_, err := strconv.Atoi(s)
// Convert rune to number: i := int(r - '0')
func main() {
	arr1 := []int{1, 1, 2, 2, 3, 4}
	//arr2 := []int{1, 3, 4}
	//arr2D := [][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}
	//arr2D := [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}
	//arrStr := []string{"5", "2", "C", "D", "+"}
	//str := "AABABBA"
	fmt.Println(deleteDuplicates(linked.NewListNode[int](arr1...)))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 1->2->3
// 1->3->4
// prev: nil, current = 1->1->2->3->3->4
// 1->1->2->3->3->4
// 1->2->3->4
func deleteDuplicates(head *linked.ListNode[int]) *linked.ListNode[int] {
	var prev *linked.ListNode[int]
	current := head
	for current != nil {
		if prev == nil {
			prev = current
			current = current.Next
		} else {
			if prev.Val == current.Val {

			}
		}
	}
	return head
}

// A -> B -> C -> D
// prev = nil
// current = head = A(addr1)
// next = B(addr2)
// A.next = prev(addr3)
// prev = A
// current = B
