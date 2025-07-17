package task_242

//242. Valid Anagram

//Given two strings s and t, return true if t is an anagram of s, and false otherwise.

//Example 1:
//Input: s = "anagram", t = "nagaram"
//Output: true
//
//Example 2:
//Input: s = "rat", t = "car"
//Output: false

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mymap := map[rune]int{}
	mymapTwo := map[rune]int{}

	for _, let := range s {
		mymap[let] = mymap[let] + 1
	}

	for _, let := range t {
		mymapTwo[let] = mymapTwo[let] + 1
	}

	for id, amount := range mymap {
		if amount != mymapTwo[id] {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mymap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		mymap[s[i]] += 1
		mymap[t[i]] -= 1
	}

	for _, amount := range mymap {
		if amount != 0 {
			return false
		}
	}
	return true
}
