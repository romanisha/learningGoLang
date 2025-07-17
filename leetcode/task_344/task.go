package task_344

/*344. Reverse String

Write a function that reverses a string. The input string is given as an array of characters s.
You must do this by modifying the input array in-place with O(1) extra memory.

Example 1:
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]*/

func reverseString(s []byte) {
	// for i := 0; i < len(s)/2; i++ {
	//     s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	// }

	i := 0
	k := len(s) - 1

	for i < k {
		s[i], s[k] = s[k], s[i]
		i++
		k--
	}
}
