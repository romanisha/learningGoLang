package task_10

import "fmt"

type Square struct {
	Side int
}

type OurType string

func Run() {
	definition()
}

// метод - это функция, которая принадлежит определенному типу или указателю на тип (receiver)
func (s Square) Perimeter() {
	fmt.Printf("%T %#v \n", s, s)
	fmt.Printf("Perimeter of the figure: %d\n", s.Side*4)
}

func (s *Square) Scale(multiplier int) { // в данном случае происходит сайд эффект и мы перезаписываем значение и оно меняется
	fmt.Printf("%T, %#v, \n", s, s) // те Pointer reciever - если нужно изменять значение изнутри функции
	s.Side *= multiplier
	fmt.Printf("%T, %#v, \n", s, s)
}

func (s Square) WrongScale(multiplier int) { // тут мы получаем копию значения, сторона не перезаписывается ТК мы используем  value reciever
	fmt.Printf("%T, %#v, \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v, \n", s, s)
}

// как вызывать методы
func definition() {
	square := Square{4}
	pSquare := &square

	square2 := Square{2}

	square.Perimeter() //
	square2.Perimeter()

	pSquare.Scale(2)

	pSquare.Perimeter() // Синтаксический сахар, позволяющийц вызывать методы наоборот и нам ничего не мешает вызват метод с велью ресивера у указателя.
	// ГО понимает, что мы передаем указатель на значение и сам разименовывает его и передает обычное значение
	square.Scale(2)
	pSquare.Perimeter()

	square.WrongScale(3)
	square.Perimeter()
}

func PusPus() {
	square1 := Square{4}
	square2 := &square1
	square3 := square1

	square2.Scale(3) // square2 = 12, также перезаписывваетсч square1 = 12 так как  square2 := &square1 это указатель, а указатель перезаписывает исходное значение
	square3.Scale(2) // square3 = 4 * 2 = 8, четверка остается так как square3 := square1 - это просто копия первого знаяения

	square1.Perimeter()
	square2.Perimeter()
	square3.Perimeter()
}
