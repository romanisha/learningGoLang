package task_1

import "fmt"

func Run() {
	name := "vasya"
	hello := fmt.Sprintf("Hello %s", name)
	_ = hello
	fmt.Println(hello)

	fmt.Printf("ТИП: %T\nValue: %v", name, name)
}
