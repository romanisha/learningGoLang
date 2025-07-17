package task_14

//14. Longest Common Prefix

/*Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.*/

func Run() {
	longestCommonPrefix()
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLenOfWord := len(strs[0])
	for i := 1; i < len(strs); i++ { //i - индекс слова в слайсе
		lenOfWord := len(strs[i]) // вычисляем длину каждого слова в слайсе
		if lenOfWord < minLenOfWord {
			minLenOfWord = lenOfWord
		}
	}

	for k := 0; k < minLenOfWord; k++ { // берем букву
		str := strs[0][k]
		for i := 1; i < len(strs); i++ {
			comp := strs[i][k]
			if comp != str {
				return strs[0][0:k] // возвращаем часть, которая совпала
			}
		}
	}
	return strs[0][:minLenOfWord]
}
