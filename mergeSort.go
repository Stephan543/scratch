package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// One list can finish before the other
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func runLLSort() {
	// Example usage:
	// Create two sorted linked lists
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	// Merge the two sorted lists
	mergedList := mergeTwoLists(list1, list2)

	// Print the merged list
	fmt.Println("Merged List:")
	printList(mergedList)
}

///////// Merge sort array

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	leftIndex, rightIndex := 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result
}

func runMergeSort() {
	// arr := []int{38, 27, 43, 3, 9, 82, 10}

	arr := []int{2, 1}
	fmt.Println("Unsorted array:", arr)

	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}

///// Quick sort

// func sortColors(nums []int) {

// 	// base case
// 	if len(nums) <= 1 {
// 		return
// 	}

// 	// Pivot
// 	p := nums[len(nums)/2]
// 	// merge sort
// 	var left, middle, right []int

// 	for _, value := range nums {
// 		if value < p {
// 			left = append(left, value)
// 		} else if value == p {
// 			middle = append(middle, value)
// 		} else {
// 			right = append(right, value)
// 		}
// 	}

// 	sortColors(left)
// 	sortColors(right)

// 	// combine results in place
// 	copy(nums, left)
// 	copy(nums[len(left):], middle)
// 	copy(nums[len(left)+len(middle):], right)
// }
