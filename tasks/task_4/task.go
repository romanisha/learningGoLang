package task_4

import "fmt"

func Run() {
	a, b := 2, 3

	// ИСПОЛЬЗОВАНИЕ ФУНКЦИИ КАК ЗНАЧЕНИЕ
	var multiplier func(x, y int) int //тут мы обьявляем функцию как переменную

	multiplier = func(x, y int) int { return x * y }
	fmt.Println(multiplier(a, b))

	divider := func(x, y int) float64 { return float64(x) / float64(y) } //а в данном случае сразу обьявляем коротким синтаксисом те мы не пишем вар и не указываем ти тк го вычислит это динамически и мы с помощью короткого синтаксиса :=
	fmt.Println(divider(b, a))

	fmt.Println(calculate(a, b, sumFunc))
	fmt.Println(calculate(a, b, subStruct))

	divier3 := CreateDivider(3)
	divider5 := CreateDivider(5)

	fmt.Println(divier3(30))
	fmt.Println(divider5(30))
}

// ФУНКЦИЯ КАК ПАРАМЕТРЫ
func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)

}
func sumFunc(x, y int) int {
	return x + y
}

func subStruct(x, y int) int {
	return x - y
}

// Возврат функций из функций
func CreateDivider(divider int) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider
	}
	return dividerFunc
}
