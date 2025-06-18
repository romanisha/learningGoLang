package task_2

import (
	"fmt"
	"unsafe"
)

func Run() {
	var num1 int64 = 15
	num2 := 15

	fmt.Println(num1 + int64(num2))

	fmt.Println(unsafe.Sizeof(uint8(num1)))
	fmt.Println(unsafe.Sizeof(uint64(num1)))
}
