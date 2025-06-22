package task_13

import "fmt"

func Run() {
	//arrays()
	//slices()
	//convertToArrayPointer()
	//passToFunction()
	//getSlice()
	//copySlice()
	deleteElements()

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
	fmt.Printf("Lenght: %d, Capacity %d\n\n", len(stringSliceLiteral), cap(stringSliceLiteral))
	//Длина слайса= кол-воэлементов, капасити - максимальная вместимость слайса до тех как заново будет переаллоцирована память

	// Создание слайса через функцию new, но это очень неудобно
	slicePointer := new([]int)
	PrintPlease(slicePointer)
	getLength(*slicePointer)

	newSlice2 := append(*slicePointer, 1, 2, 3)
	PrintPlease(newSlice2)
	getLength(newSlice2)

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

// получение слайса из массива, получение слайса из слайса
// в го существует специальная операция  получения слайса, слайс можем получать как на основе массива, так и на основе др. слайса
// у нас есть массив и мы хотим слайс на основе этого масива
func getSlice() {

	intArr := [...]int{1, 2, 3, 4, 5}
	PrintPlease(intArr)

	//Реслайсинг (1 способ получение слацса из массива)
	intSlice := intArr[1:3] //мы создаем слайс от элемента с индексом 1 (включая) до элемента с индексом 4 (не включая)
	PrintPlease(intSlice)
	getLength(intSlice) // капасити в данном случае считается от выбранного элемента до конца ИСХОДНОГО массива

	fullSlice := intArr[:] // такая запись берет ВСЕ значения int[0:5]
	PrintPlease(fullSlice)
	getLength(fullSlice)

	sliceFromSlice := fullSlice[0:3] // слайсот другого слайса
	PrintPlease(sliceFromSlice)
	getLength(sliceFromSlice)

	//все слайсы ссылаются на исходный массив, поэтому изменив значение в массиве, мы также их увидим во всех последующих от него слайсов
	intArr[2] = 500
	PrintPlease(intArr)
	PrintPlease(intSlice)
	PrintPlease(fullSlice)
	PrintPlease(sliceFromSlice)
}

// КОПИРОВАНИЕ СЛАЙСОВ
func copySlice() {
	destination := make([]string, 0, 2)
	sourse := []string{"Vasya", "Petya", "Katya"}

	makeCopy(destination, sourse) // в данном случае копирование не произойдет
	PrintPlease(destination)      //и нас дестинейшен не изменится
	fmt.Printf("Lenght: %d, Capacity %d\n\n", len(destination), cap(destination))

	destination = make([]string, 2, 3)
	makeCopy(destination, sourse) // копируется столько значений, сколько у нашего слайса длины, у нас 2 - скорировалось 2
	PrintPlease(destination)      // капасити в данном случае вообще не при чем
	fmt.Printf("Lenght: %d, Capacity %d\n\n", len(destination), cap(destination))

	destination = make([]string, len(sourse)) // тк нам нужно скопировать весь source, то выставляем длину source
	makeCopy(destination, sourse)
	PrintPlease(destination)
	fmt.Printf("Lenght: %d, Capacity %d\n\n", len(destination), cap(destination))

	var defaultSlice []string // ниловый слайс
	PrintPlease(defaultSlice)
	fmt.Printf("Lenght: %d, Capacity %d\n\n", len(defaultSlice), cap(defaultSlice))

	makeCopy(defaultSlice, sourse) //после копирования ниловый слайс таким и останется
	PrintPlease(defaultSlice)
	fmt.Printf("Lenght: %d, Capacity %d\n\n", len(defaultSlice), cap(defaultSlice))

	//для копирования данных в дефолтный слайс мы можем исполтзовать связку мейка и аппенда
	rightSlice := append(make([]string, 0, len(sourse)), sourse...)
	PrintPlease(rightSlice)
	fmt.Printf("Lenght: %d, Capacity %d\n\n", len(rightSlice), cap(rightSlice))
}
func convertToArrayPointer() {
	initialSlice := []int{1, 2}
	PrintPlease(initialSlice)
	getLength(initialSlice)

	intArray := (*[2]int)(initialSlice) //конвертация слайса в указатель на массив.
	// Чтобы это сработало нужно, чтобы длина массива совпадала с длиной слайса
	PrintPlease(intArray)
	fmt.Printf("Lenght: %d, Capacity %d\n\n", len(intArray), cap(intArray))
}

// Удаление элемента из слайса
func deleteElements() {
	slice := []int{1, 2, 3, 4, 5}
	i := 2
	PrintPlease(slice)
	getLength(slice)

	//1ый способ удаления через append, но он ломает исходный слайс
	withAppend := append(slice[:i], slice[i+1:]...) //  те мы берем до 2ого элемента и после третьего до конца
	PrintPlease(withAppend)
	getLength(withAppend)
	PrintPlease(slice) // ломает исходный слайс,чтобы этого не было лучше переприсваивать исходный слайс (не использовать новый слвй слайс )

	slice = []int{1, 2, 3, 4, 5}

	withCopy := slice[:i+copy(slice[i:], slice[i+1:])] //ломает исходный слайс, чтобы этого не было меняем withCopy на slice и работаем с одним слайсом
	//slice = slice[:i+copy(slice[i:], slice[i+1:])]
	PrintPlease(withCopy)
	fmt.Println(slice)
}

// передача слайса в функцию. В ГО все передается по значению
func passToFunction() {
	initialSlice := []int{1, 2}
	PrintPlease(initialSlice)
	getLength(initialSlice)

	changeValue(initialSlice) //тк слайс содержит указатель на массив, значит изменяя значение в нашес слайсе
	// то мы меняем значение  и в исходном масиве
	PrintPlease(initialSlice)
	getLength(initialSlice)

	//initialSlice = append(initialSlice, 3) // добавляем элементы слайса и исходный сдайс перезаписывается
	newSlice := append(initialSlice, 3)
	PrintPlease(newSlice)
	getLength(newSlice)

	newSlice = appendValue(newSlice)
	PrintPlease(newSlice)
	getLength(newSlice)
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

func appendValue(slice []int) []int {
	slice = append(slice, 4, 5)
	PrintPlease(slice)
	getLength(slice)

	return slice
}

func changeValue(slice []int) {
	slice[1] = 15
}

func PrintPlease(t any) {
	fmt.Printf("ТИП: %T, Значение %#v\n", t, t)
}

func ChangeArray(arr [3]int) [3]int {
	arr[2] = 3
	return arr
}

func GetIndex(i, v any) {
	fmt.Printf("Index: %d, Value: %v\n", i, v)
}

func getLength(l []int) {
	fmt.Printf("Lenght: %d, Capacity %d\n\n", len(l), cap(l))
}

func makeCopy(d, s []string) {
	fmt.Printf("Copied %v\n", copy(d, s))
}
