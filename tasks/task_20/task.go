package task_20

import (
	"fmt"
)

func Run() {
	//showSum()
	//ShowContains()
	//ShowAny()
	//unionInterfaceAndType()
	typeApproximation()
}

type Number interface {
	~int64 | float64 // мы создаем интерфейс ТИПОВ ( БЕЗ МЕТОДОВ) чтобы использовать в дженериках в качестве констрейнта
	// если добавляем перед инт тильду (~), это значит мы используем приближение типа( те используюем типы, которые основанвы на ИНТ64)
}

type CustomInt int64 // кастомный тип, создаем когда хотим расширить какое-либо поведения стандартных типов

func (ci CustomInt) IsPositive() bool {
	return ci > 0
}

type Numbers[T Number] []T //слайс элементов типа Т

//Дженерики - данная фича позволяет выполнять обощение типов
// в ГО два вида  обобщения типов: 1) для функции 2) обощенные типы

func showSum() {

	floats := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	ints := []int64{1, 2, 3, 4, 5}

	fmt.Println(sum(floats))
	fmt.Println(sum[int64](ints)) //можно явно указать тип

}

func ShowContains() {
	type Person struct {
		name     string
		age      int
		jobTitle string
	}
	ints := []int64{1, 2, 3, 4, 5}
	fmt.Println("Ints:", contains(ints, 4))

	strings := []string{"a", "b", "c", "d"}
	fmt.Println("Strings:", contains(strings, "d"))
	fmt.Println("Strings:", contains(strings, "e"))

	people := []Person{
		{"Vasya",
			25,
			"frontEnder",
		},
		{"Dasha",
			18,
			"Student"},
		{
			"Pasha",
			32,
			"Dolbayob",
		},
	}

	fmt.Println("Structs:", contains(people, Person{
		name:     "Vasya",
		age:      21,
		jobTitle: "frontEnder",
	}))

	fmt.Println("Structs:", contains(people, Person{
		name:     "Dasha",
		age:      18,
		jobTitle: "Student",
	}))
}

func ShowAny() {
	// встроенные интерфейсы для дженериков: comparable и any
	show(1, 2, 3)
	show("test1", "test2", "test3")
	show([]int{1, 2, 3}, []int{4, 5, 6})
	show(map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	})
	show(interface{}(1), interface{}("string"), any(struct {
		name string
	}{name: "Vasya"}))

}

func unionInterfaceAndType() {
	var ints Numbers[int64]                        // создаем обобщеный тип намберс, обязательно указываем тип, который хотим туда класть
	ints = append(ints, []int64{1, 2, 3, 4, 5}...) // добавляем слайс

	floats := Numbers[float64]{1.0, 2, 3, 5, 5}

	fmt.Println(sumUnionInterface(ints))
	fmt.Println(sumUnionInterface(floats))
}

// создание дженериков
// в квадратных скобках Type parameters, мы создаем тип и указываем какого базовгово типа может быть данный тип. Если понимаем, что типов несколько, то через | указываем неск
func sum[V int64 | float64](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func contains[V comparable](sliceElements []V, findValue V) bool { //тип comparable - все типы, которые можно срввнить( те все, кроме мап, слайсов, массивов  с некомперебл элементы)
	for _, val := range sliceElements {
		if val == findValue {
			return true
		}
	}
	return false
}

func typeApproximation() {
	customInts := []CustomInt{1, 2, 3, 5, 6} // пользуемся нашим кастомным типом и хотим его передать в функцию, который используют джнерики
	castedInts := make([]int64, len(customInts))

	for idx, val := range customInts {
		castedInts[idx] = int64(val)
	}
	fmt.Println(sumUnionInterface(customInts)) //не сработает, потому что передает слайс кастомных типов, которые не соответсвуют типам в функции
	// чтобы сработало можно либо преобразорвать в нужный нам тип посредством создания новоцй переменной castedInts
	//либо исползоать "приближение типов" см строку 16
	fmt.Println(sumUnionInterface(castedInts))

}

func show[T any](entities ...T) { // ...T это упаковка в слайс, те мы передаем не слайс целиком, а сразу по значениям
	fmt.Println(entities)
}

func sumUnionInterface[V Number](numbers []V) V { // в качестве констрейнта указан намбер, который принимает инт64 иили флоат64
	var sum V
	for _, num := range numbers {
		sum += num
	}
	return sum

}
