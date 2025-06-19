package task_13

import "fmt"

func Run() {
	arrays()
	slices()
}

type Person struct {
	name string
	age  int
}

// Массивы - тип данных, который  может содержать фиксированное кол-во элементов определенного типа

// Создание массива
func arrays() {
	var intArr [3]int // для создания массива обьявляем переменную с квдр. скобками
	//в квадратных скобках - кол-во элементов массива
	//те мы создаем массив из ТРЕХ элементов
	PrintPlease(intArr)

	intArr[0] = 5
	intArr[1] = 6
	PrintPlease(intArr)

	people := [2]Person{ //массив содержит два элемента типа Персон
		{
			name: "Bob",
			age:  10,
		},
		{
			name: "Alice",
			age:  17,
		},
	}
	PrintPlease(people)

	// у нас также есть возможность мсоздать массив, где не нужно указввать кол-во элементов
	//го вычеслит это кол-во автоматически

	// создаем переменнуюи присваиваем ей массив строк
	stringsArr := [...]string{"First", "Second", "Third", "Fourth"}
	PrintPlease(stringsArr)

	//вызов встроенных функций длина (len) и вместимость (cap). Длина массиваеизменяема
	// для массива длина и вместимость  ВСЕГДА будет одинаковая и она будет равнятся той длине,которую мы указываем при обьявлении массива

	fmt.Printf("Lenght: %d, Capacity %d\n", len(stringsArr), cap(stringsArr))
	//GetIndex(stringsArr)

	// Интеграции по массиву
	// итерации с помощью цикла for
	for i := 0; i < len(stringsArr); i++ {
		GetIndex(i, stringsArr[i])
	}

	// итерации с помощью цикла for range, нам не нужно знать размер цикла
	//в первую переменную всегда приходит индекс, на каждыой итерации индекс будет увеличиваться а 1
	// второй элемент - значение
	for inx, value := range stringsArr {
		GetIndex(inx, value)
	}

	// Допустим нам при переборе массива индекс не нужен
	for _, value := range stringsArr {
		fmt.Printf("Value: %s\n", value)
	}

	newIntArr := ChangeArray(intArr)
	PrintPlease(intArr)
	PrintPlease(newIntArr)

}

func slices() {

}

func PrintPlease(t any) {
	fmt.Printf("Тип: %T, Значение %#v\n", t, t)
}

func GetIndex(i, v any) {
	fmt.Printf("Index: %d, Value: %s\n", i, v)
}

func ChangeArray(arr [3]int) [3]int {
	arr[2] = 3
	return arr

}
