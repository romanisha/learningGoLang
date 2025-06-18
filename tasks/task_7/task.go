package task_7

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	min = 1
	max = 5
)

func Run() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := r.Intn(max-min) + min
	SwitchTry(a)
}

// Конструкция Switch case - заменя нескольцих IF и ELSE

func SwitchTry(value int) {

	//if example
	if value == 1 {
		fmt.Println("number 1")
	} else if value == 2 || value == 3 {
		fmt.Println("number 2 or 3")
	} else if value == GetFour() {
		fmt.Println("number is 4")
	} else {
		fmt.Println("smth strange")
	}

	//base switch
	switch value {
	case 1:
		fmt.Println("number 1")
	case 2, 3:
		fmt.Println("number 2 or 3")
	case GetFour():
		fmt.Println("number is 4")
	default:
		fmt.Println("smth strange")
	}
	//switch с локальным обозначением переменной
	switch num := rand.Intn(max-min) + min; num {
	case 1:
		fmt.Println("number 1")
	case 2:
		fmt.Println("number 2 or 3")
	case GetFour():
		fmt.Println("number is 4")
		fallthrough //говорит, что нужнг проваливаться дальше, даже если условие уже выполнено.Т е получая 4ку мы выполняем код и проваливаемся дальше  к 10 и вынуждены выполнить ее
	case 10:
		fmt.Println("this is 10")
	default:
		fmt.Println("smth strange")
	}

	//switch без условий
	switch {
	case value > 2:
		fmt.Printf("Value %d greater than 2\n", value)
	case value < 2:
		fmt.Printf("Value %d less than 2\n", value)
	default:
		fmt.Println("Value is 2")
	}
}

func GetFour() int {
	return 4
}
