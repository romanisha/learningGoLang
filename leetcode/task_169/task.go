package task_169

/*169. Majority Element

Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

Example 1:
Input: nums = [3,2,3]
Output: 3

Example 2:
Input: nums = [2,2,1,1,1,2,2]
Output: 2*/

func Run() {

}

func majorityElement(nums []int) int {
	myMap := make(map[int]int)
	for _, num := range nums {
		myMap[num] = myMap[num] + 1
	}
	for k, v := range myMap {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
