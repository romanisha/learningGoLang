package task_136

import "fmt"

func Run() {

	nums := []int{1, 1, 2}
	fmt.Println(singleNumber(nums))
}

/*136. Single Number
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.

		Example 1:
Input: nums = [2,2,1]
Output: 1

Example 2:
Input: nums = [4,1,2,1,2]
Output: 4

Example 3:
Input: nums = [1]
Output: 1 */

func singleNumber(nums []int) int {

	myMap := map[int]int{}

	for _, num := range nums {
		myMap[num] = myMap[num] + 1
	}

	for key, amount := range myMap {
		if amount == 1 {
			return key
		}
	}
	return -1
}
