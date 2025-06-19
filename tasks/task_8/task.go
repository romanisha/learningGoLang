package task_8

import "fmt"

func Run() {
	//УКАЗАТЕЛИ
	// ДЕФОЛТНОЕ ЗНАЧЕНИЕ

	var intPointer *int //указатель на тип инт
	fmt.Printf("%T %#v\n", intPointer, intPointer)

	//получение non nil указателей
	var a int64 = 7
	fmt.Printf("%T %#v\n", a, a)

	var PointerA *int64 = &a                                   // амперсанд а - указатель на а
	fmt.Printf("%T %#v %#v\n ", PointerA, PointerA, *PointerA) //когда мы вытаскиваем велью из указателя, то мы получаем просто адрес (второе значение)
	// третье значение - мы получаем значение из памяти с помощью операции разименования *PointerA
	// Если указатель будет ссылаться на нил, то нужно перед разименованием проверять на нил, иначе будет паника

	var newPointer = new(int64)
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer) //Получаем сначала ссылку, потом реальное значение
	*newPointer = 2                                                  // меняем значение указателя
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)

	// как использовать показатели?
	//side effect
	num := 3
	num = square(num)
	fmt.Println(num)

	SquarePointer(&num)
	fmt.Println(num)

	//empty value признак пустого значени
	// проверяем есть ли у человека кошелек с помощью указателя на int тк у него есть вариант иметь значение nil
	// те может бфть сумма на счету 0, но это не значит, что кошелька нет
	var wallet1 *int // с помощью указателя идентифицируем nil
	fmt.Println(hasWallet(wallet1))

	wallet2 := 0
	fmt.Println(hasWallet(&wallet2))

	wallet3 := 1000
	fmt.Println(hasWallet(&wallet3))
}

func square(num int) int {
	return num * num
}

func SquarePointer(num *int) {
	*num *= *num // равно *num = *num * *num
}

func hasWallet(money *int) bool {
	return money != nil
}
