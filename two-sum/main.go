/*
 * https://leetcode.com/problems/two-sum/
 */

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			num1 := nums[i]
			num2 := nums[j]

			if num1+num2 == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)

	res = twoSum([]int{3, 2, 4}, 6)
	fmt.Println(res)

	res = twoSum([]int{3, 3}, 6)
	fmt.Println(res)
}
