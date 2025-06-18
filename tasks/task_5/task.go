package task_5

import "fmt"

func Run() {
	age := 20

	if age < 18 { //Условный оператор  IF
		//fmt.Println("you r too young")
	}
	//Короткий синтаксис
	if isChild := isChildren(age); isChild == true { // тут мы объявляем переменную isChild и присваеваем ей значение, которое является результатом выполнения функции
		// после ; идет плок проверки условия
		//fmt.Println("yo yo molod suka")
		//fmt.Println(isChild)
	}

	//else
	if age < 18 {
		//fmt.Println("you r too young")
	} else {
		fmt.Println("vse norm")
	}
	//Логическое умножение && которое выполняет условие и то и то
	if age > 6 && age < 18 {
		fmt.Println("pupil")
	}
	//Логическое сложение || которое выполняет хотя бы ОДНО условие.
	if age == 14 || age == 20 || age == 40 {
		fmt.Println("get passport")
	}
}

func isChildren(age int) bool {
	return age < 18
}
