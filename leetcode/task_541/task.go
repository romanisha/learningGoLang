package task_541

//541. Reverse String II

//Given a string s and an integer k, reverse the first k characters for every 2k characters counting from the start of the string.
//If there are fewer than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and leave the other as original.

//Example 1:
//Input: s = "abcdefg", k = 2
//Output: "bacdfeg"

//Example 2:
//Input: s = "abcd", k = 2
//Output: "bacd"

func reverseStr(s string, k int) string {
	if k == 0 || k == 1 {
		return s
	}

	//var newString []byte
	var newString string
	tf := true
	for i := 0; i < len(s); i = i + k {

		firstIndex := i
		lastIndex := i + k

		if lastIndex > len(s) {
			lastIndex = len(s)
		}

		if tf {
			newString = newString + reverse(s[firstIndex:lastIndex])
			tf = false
		} else {
			newString = newString + s[firstIndex:lastIndex]
			tf = true
		}

	}

	return newString
}

func reverse(r string) string {
	var newR string
	for i := len(r) - 1; i >= 0; i-- {

		newR = newR + string(r[i])
	}

	return newR
}

// GOOD AND EASY SOLUTION
func reverseStr(s string, k int) string {
	conv := []rune(s)
	n := len(conv)

	for i := 0; i < n; i += 2 * k {
		l, r := i, min(i+k-1, n-1)
		for l < r {
			conv[l], conv[r] = conv[r], conv[l]
			l++
			r--
		}
	}
	return string(conv)
}
