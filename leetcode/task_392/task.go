package task_392

/*Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:
Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:
Input: s = "axc", t = "ahbgdc"
Output: false*/

func isSubsequence(s string, t string) bool {

	i, k := 0, 0
	for i < len(s) && k < len(t) {
		if s[i] == t[k] {
			i++
			k++
		} else {
			k++
		}
	}

	if i == len(s) {
		return true
	}
	return false

}
