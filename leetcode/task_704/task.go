package task_704

/*
704. Binary Search

Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	halfOfSlice := len(nums) / 2
	firstIndex := 0
	lastIndex := len(nums)
	for {
		if nums[halfOfSlice] == target {
			return halfOfSlice
		}

		if firstIndex == halfOfSlice || lastIndex == halfOfSlice {
			return -1
		}

		if nums[halfOfSlice] < target {
			firstIndex = halfOfSlice
		} else if nums[halfOfSlice] > target {
			lastIndex = halfOfSlice
		}

		halfOfSlice = (lastIndex + firstIndex) / 2
	}
}
