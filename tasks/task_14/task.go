package task_14

import "fmt"

type User struct {
	Id      int64
	Name    string
	Surname string
}

func Run() {
	createMap()
}

//Мапы -тип данных, созданных на основе хэш таблицы
//хэш таблица - структура данных со значениями в виде пар: ключа и значения
// ключ должен быть comparable
// несравниваемые типы: слайсы, мапы, функции

func createMap() {
	//default value
	var defaultMap map[int64]string //в квадратных скобках - тип ключа, за скобками тип значения
	printType(defaultMap)
	getLength(defaultMap)

	//map by make
	mapByMake := make(map[string]string)
	printType(mapByMake)

	//map by make with cup
	mapByMakeWithCup := make(map[int64]string, 3) // указываем кол-во элементов, которые хотим разместить в нашей мапе
	printType(mapByMakeWithCup)

	//map by literal
	mapByLiteral := map[string]int{"Vasya": 18, "Dima": 20} //ключ: значение
	printType(mapByLiteral)
	getLength(mapByLiteral)

	//map by new
	mapWithNew := *new(map[string]string) // через нью мы всегда вызываем указатель на тип, поэтому мы разименовываем сразу
	printType(mapWithNew)

	//insert value
	mapByMake["First"] = "Vasya" //в квадратных скобках создаем ключ мапы и присваиваем ей знаяение Вася
	printType(mapByMake)
	getLength(mapByMake)
	//update value
	mapByMake["First"] = "Petya"
	printType(mapByMake)
	getLength(mapByMake)

	// get map value
	fmt.Println(mapByLiteral["Vasya"])
	//get map default value
	fmt.Println(mapByLiteral["second"]) //пытаемся вызвать по несуществующему ключу и получаем 0, те получаем дефолтное значение типа

	//check value existence мы можем выяснить действительно ли значение в мапе 0, или оно просто не существует вообще
	value, ok := mapByLiteral["second"]
	fmt.Printf("Value: %v Is Exists: %t\n", value, ok)
	//isMapExist(mapByMake["second"])

	//delete value, функция удаления данных по ключу
	delete(mapByMake, "First")
	printType(mapByMake)

	//map iteration
	mapForIteration := map[string]int{"First": 1, "Second": 2, "Third": 3, "Fourth": 4}

	for key, val := range mapForIteration {
		fmt.Printf("Key: %v, Val: %v\n", key, val) // на последовательность в мапах полагаться нельзя, при каждом запуске цикла она может быть разная
	}

	//unique value - использование мапы как сета, те создаем набор уникальных сущностей
	users := []User{
		{
			Id:      1,
			Name:    "Vasya",
			Surname: "Pupkin",
		},
		{
			Id:      45,
			Name:    "Petya",
			Surname: "Snow",
		},
		{
			Id:      57,
			Name:    "John",
			Surname: "Smith",
		},
		{
			Id:      45,
			Name:    "Petya",
			Surname: "Snow",
		},
	}
	UsersUnique := map[int64]struct{}{} //в качестве значения будем использовать пустую структуру
	for _, user := range users {
		if _, ok := UsersUnique[user.Id]; !ok { // // Достаем из мапы пользователя (user) по айдишнику user.id - получаем (_, ок)
			UsersUnique[user.Id] = struct{}{} // Кладем в мапу элемент из слайса user по ключу равному user.Id (то есть его айдишник)
		}
	}
	printType(UsersUnique)

	// Поиск по значений
	usersMap := map[int64]User{}
	for _, user := range users {
		if _, ok := usersMap[user.Id]; !ok {
			usersMap[user.Id] = user
		}
	}
	fmt.Println(findIdnSlice(57, users))
	//использование мапы для поиска
	fmt.Println(findInMap(34, usersMap))

}

// быстрый поиск значения с помощью мап
// данный способ будет работать долго тк у нас может быть пользователь в конце списка и придется циклом пробегаться по всему циклу
// сложность выполнения (О)n
func findIdnSlice(id int64, users []User) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func findInMap(id int64, usersMap map[int64]User) *User {
	if user, ok := usersMap[id]; ok {
		return &user
	}
	return nil
}
func printType[K comparable, T any](t map[K]T) {
	fmt.Printf("Type: %T, Value: %v\n", t, t)
}

func getLength[K comparable, T any](t map[K]T) {
	fmt.Printf("Length: %d\n\n", len(t))
}

//func isMapExist[K comparable, T any](t map[K]T) {
//	fmt.Printf("Value: %v Is Exist: %t\n\n", t, t)
//}
