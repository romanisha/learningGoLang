package task_9

import (
	"fmt"
	"time"
)

// кастомные типы -тип, который содаем сами, указываем базовый тип значения (строка, инт и тд)
type OurString string
type OurInt int

// структура - набор полей, обьявляется с помощью ключевого слова type
type Person struct {
	name string
	age  int
	town string
}

func Run() {

	var customString OurString
	fmt.Printf("%T %#v \n", customString, customString)
	customString = "Hello World"
	fmt.Printf("%T %#v \n", customString, customString)

	customInt := OurInt(10) //создание переменной с помощью короткого синтаксиса
	fmt.Printf("%T %#v \n", customInt, customInt)
	fmt.Println(int(customInt)) // если мы хотим привести наш кастомный тип customInt к int (к своему базовому типу)

	//Инициализация структуры с дефолтным значением - те значения которые присвоины типам на базовом урлвне
	var John Person
	fmt.Printf("%T %#v \n", John, John)

	John = Person{}
	fmt.Printf("%T %#v \n", John, John)

	//доступ к полям структуры, полуячение их значений и изменение значений
	// к полям структыры мы обращаемся через точку. Указываем переменную  и название поля John.name
	John.name = "John Doe"
	John.age = 10
	John.town = "Moscow"
	fmt.Println(John)

	// инициализация переменной структуры в фигурных скобках
	Bob := Person{
		name: "Bob",
		age:  20,
		town: "Moscow",
	}
	fmt.Println(Bob)

	Vlad := Person{"Vlad", 17, "Moscow"}
	fmt.Println(Vlad)

	// Получение полей у указателя на структуру
	pVlad := &Vlad            //указатель на переменную Влад
	fmt.Println((*pVlad).age) // мы снрачала разименовываем показатель, чтобы получать значение поля.
	fmt.Println(pVlad.age)    // синтаксический сахар, эта и верхняя строка равнозначны

	//создание только указателя на структуру
	pAnastasiia := &Person{"Anastasiia", 25, "Moscow"}
	fmt.Println(pAnastasiia)

	//Анонимные структуры
	AnonStructure := struct {
		name, lastName, birthday string
	}{
		name:     "noName",
		lastName: "noLastName",
		birthday: time.Now().String(), // через точку вызываем метод
	}
	fmt.Println(AnonStructure)
}
