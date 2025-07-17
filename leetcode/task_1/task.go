package task_1

/*1. TwoSum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.*/

/*
Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]
*/

func Run() {
	twoSum()
}

func twoSum(nums []int, target int) []int {
	tt := map[int]int{}

	for i := 0; i < len(nums); i++ {
		v, ok := tt[nums[i]]

		if ok {
			return []int{v, i}
		}

		tt[target-nums[i]] = i
	}

	return []int{}
}

// func twoSum(nums []int, target int) []int {
//     for i := 0; i <= len(nums) - 1; i++ {
//         // тут мы должны начат складывать и сравнивать элементы
//         for y := i + 1; y <= len(nums) - 1; y++ {
//             res:= nums[i] + nums[y]
//             if res == target {
//                    return []int{i, y}
//             }
//         }
//     }
// return []int{}
// }
