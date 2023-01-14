/*
 * https://leetcode.com/problems/shuffle-the-array
 */

package main

import "fmt"

// assume `nums` is ordered as [x1,x2,...,xn,y1,y2,...,yn]
// also n * 2 = len(nums)
// returned array must be in the form of `[x1,y1,x2,y2,...,xn,yn]`
func shuffle(nums []int, n int) []int {
	nums1 := nums[:n]
	nums2 := nums[n:]
	res := make([]int, n*2)

	pos, i := 0, 0
	for i < n {
		res[pos] = nums1[i]
		res[pos+1] = nums2[i]

		pos += 2
		i++
	}

	return res
}

func main() {
	res := shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
	fmt.Println(res)
}
