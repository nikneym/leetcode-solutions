/*
 * https://leetcode.com/problems/median-of-two-sorted-arrays
 */

package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	nums := make([]int, total)
	i, j, pos := 0, 0, 0

	for i < len(nums1) && j < len(nums2) {
		n1 := nums1[i]
		n2 := nums2[j]

		if n1 < n2 {
			nums[pos] = n1
			i++
		} else {
			nums[pos] = n2
			j++
		}

		pos++
	}

	for i < len(nums1) {
		nums[pos] = nums1[i]
		i++
		pos++
	}

	for j < len(nums2) {
		nums[pos] = nums2[j]
		j++
		pos++
	}

	isOdd := total%2 == 1

	if isOdd {
		n := nums[total/2]
		return float64(n)
	}

	// length is even
	n1 := float64(nums[total/2-1])
	n2 := float64(nums[total/2])
	return (n1 + n2) / 2
}

func main() {
	res := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println(res)

	res = findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(res)
}
