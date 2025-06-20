package task_13

import "fmt"

func Run() {
	arrays()
	slices()
	convertToArrayPointer()
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

//Слайсы (срезы) - тип данных, который может содержать любое количество элементов.
//Их называют динамическими массивами

func slices() {
	//создание слайса, кол-во элементов мы не указываем
	//слайс с дефолтным значением
	var defaultSlice []int
	PrintPlease(defaultSlice) // дефолтное значение слайса - это нил
	getLength(defaultSlice)

	// второй вариант создание слайсов - использование литералов
	stringSliceLiteral := []string{"First", "Second"}
	PrintPlease(stringSliceLiteral)
	fmt.Printf("Lenght: %d, Capacity %d\n", len(stringSliceLiteral), cap(stringSliceLiteral))
	//Длина слайса= кол-воэлементов, капасити - максимальная вместимость слайса до тех как заново будет переаллоцирована память

	//Создание слайса с помощью make
	sliceByMake := make([]int, 0, 5) // длина слайса - 0, 5 - вместимост слайса
	PrintPlease(sliceByMake)
	getLength(sliceByMake)

	// Слайс не хранит данные напрямую, это указатель на массив
	// при создании  слайса под капотом Го создает массив (те слайс ссылается на массив)

	// функция APPEND - добавляет элемент в наш слайс
	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5)
	PrintPlease(sliceByMake)
	getLength(sliceByMake)
	// если мы попытаемся добавить еще один элемент и го поймет, что кол-во элементов станет больше капасити
	// то в этот момент ГО выделит в памяти новый массив, у которого емкость будет в 2 раза больше, чем у предыдущего слайса
	// и наш новый слацс будет ссылаться на новый массив в памяти
	//капасити в слайсах ВСЕГД[А увеличивается в 2 раза
	sliceByMake = append(sliceByMake, 6)
	PrintPlease(sliceByMake)
	getLength(sliceByMake)

	//for range для слайсов
	for ind, val := range sliceByMake {
		GetIndex(ind, val)
	}

	showAllElements(1, 2, 3)
	showAllElements(1, 2, 3, 4, 5, 6, 7)

	//представим  такой кейс, что у нас есть какой-то слайс и у него есть сколько-то элементов
	// И мы быхотели вывести все его элементы в нашей функции
	// но наша функция (шоуАллЭлементс) принимает все элементы перечисленные через запятую, а не слайс целиком
	// для этого в ГО есть специальный синтаксис, когда мы можем разложить наш массив или слайс на неск элементов
	//и отправить их в качестве параметров функции
	firstSlice := []int{5, 6, 7, 8}
	secondSlice := []int{9, 3, 2, 1}
	showAllElements(firstSlice...) // для этого указываем наш слайс и добавляем 3 точки, такая запись равна перечислению элементов при вызове функции

	newSlice := append(firstSlice, secondSlice...) // добавление элементов слайса
	PrintPlease(newSlice)

}

func convertToArrayPointer() {
	initialSlice := []int{1, 2}
	PrintPlease(initialSlice)
	getLength(initialSlice)

	intArray := (*[2]int)(initialSlice) //конвертация слайса в указатель на массив.
	// Чтобы это сработало нужно, чтобы длина массива совпадала с длиной слайса
	PrintPlease(intArray)
	fmt.Printf("Lenght: %d, Capacity %d\n", intArray, intArray)
}

// Функции с неограниченным числом параметров ( Variadic function)
// чтобы обьявить такую функцию надо понимать, что такое слайсы и как они работают
func showAllElements(values ...int) {
	//перед типом аргумента указываем 3 точки, это значит, что данная функция может
	//принять сколько угодно параметров укаданного типа
	// А все аргументы функции будут сложены в слайс, который называется у нас values
	for _, val := range values {
		fmt.Println("Value:", val)
	}
	fmt.Println()
}

func PrintPlease(t any) {
	fmt.Printf("ТИП: %T, Значение %#v\n", t, t)
}

func GetIndex(i, v any) {
	fmt.Printf("Index: %d, Value: %v\n", i, v)
}

func ChangeArray(arr [3]int) [3]int {
	arr[2] = 3
	return arr
}

func getLength(l []int) {
	fmt.Printf("Lenght: %d, Capacity %d\n\n", len(l), cap(l))
}
