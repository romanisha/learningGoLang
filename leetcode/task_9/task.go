package task_9

func isPalindrome(x int) bool {
	Slice := []int{x}
	y := len(Slice) - 1

	for i := 0; i < (len(Slice))/2; i++ {
		if i == y {
			y--
		} else {
			return false
		}
	}
	return true
}
