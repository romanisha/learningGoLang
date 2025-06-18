package task_3

import "fmt"

func Run() {
	a, b := 2, 3
	sum, mult := NamedSumAndMultiply(a, b) // сам и мулт никак не связаны с аргументами функциями
	fmt.Println(sum, mult)

}

func Greet() {
	fmt.Println("Hello my friend")
}
func PersonGreet(name string) {
	fmt.Printf("Privetik, %s!\n", name)
}
func NameSurDreetings(name, surname string) {
	fmt.Printf("hihihi %s %s!!!\n", name, surname)
}

func Sum(a, b int) int {
	sum := a + b

	return sum
}

func Minus(a, b int) int {
	minus := a - b

	return minus
}

// 2a + b - 2b
func Uravnenie(a, b int) int {
	first := Sum(2*a, b)
	second := Minus(first, 2*b)

	return second
}

func SumAndMultiply(a, b int) (int, int) { //вторые скобки - результат который ожидаем т.е возвращаемое значение
	return a + b, a * b // Достаточно простооставить ретерн
}

// Именнованные возвращаемые значения
func NamedSumAndMultiply(a, b int) (sum int16, mult int16) {
	sum = int16(a + b) // Так как выше (во вторых скобках) мы указали наименования возвращаемых значение, то в данной строчке уже присваивать ( sum := )  не нужно!
	mult = int16(a * b)

	return // или return sum, mult ( в данном случае также можно не указывать  ожидаемый результат, тк мы указали его в 46 строке
}
