/*
 * https://leetcode.com/problems/happy-number/
 */

package main

import "fmt"

func calculate(n int) int {
	var digits []int
	temp := n
	for temp > 0 {
		digits = append(digits, temp%10)
		temp /= 10
	}

	var sum int = 0
	for _, digit := range digits {
		sum += digit * digit
	}

	return sum
}

func isHappy(n int) bool {
	temp := n
	var encountered []int

	for {
		temp = calculate(temp)
		if temp == 1 {
			return true
		}

		for _, digit := range encountered {
			if temp == digit {
				return false
			}
		}

		encountered = append(encountered, temp)
	}
}

func main() {
	fmt.Println(isHappy(100))
}
