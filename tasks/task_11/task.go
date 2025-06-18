package task_11

import "fmt"

//  Интерфейс - специальный тип в го, представляющий ииз себя набор сигнатур методов
//обьявляем интерфыейс

func Run() {
	interfaceValues()
	howToUseInterface()

}

type Runner interface {
	Run() string // данная строка обозначает метод интерфейса, те говорит, что у данного интерфейса должен быть реализован этот метод
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

type Docker interface {
	Swimmer
	Flyer
	Runner
}
type EmptyInterface interface{}

type Duck struct {
	Name, Surname string
}

func (d Duck) Run() string {
	return "Утка бегает"
}

func (h Human) Run() string {
	return fmt.Sprintf("Человек %s бегает", h.Name)
}

func (h Human) WriteCode() {
	fmt.Println("Челикс пишет код")
}
func (d Duck) Fly() {
	fmt.Println("Уточка может летать")
}
func (d Duck) Swim() string {
	fmt.Println("Duck is swimming!")
	return "Duck is swimming!"
}

// Существуют значения ИНТЕРФЕЙСТНОГО и КОНКРЕТНОГО типа
// создаем значение интерфейсного типа
// интерфейсные значения хранят в себе конкретный тип и значение
// в данном примере мы имеем дефолтные значения, тк мы ничего не присвоили
func interfaceValues() {
	var swimmer Swimmer
	printPlease(swimmer)

	if swimmer == nil {
		fmt.Println("swimmer is nil")
	}

	var NoValueSwimmer *Human
	printPlease(NoValueSwimmer)

	if NoValueSwimmer == nil {
		fmt.Println("NoValueSwimmer is nil")
	}

	swimmer = NoValueSwimmer // присвоение значения интерфейсному типу, то есть когда мы помещаем переменную
	//с типом указатель на Хьюман (42 строка) в наше интерфыейсное значение свиммер,
	//то у нашего интерфейсного значения появляются знание
	// о конкретном типе - указатель на хьюман, и о конкретном значении, мы видим, что значение NIl
	//если мы сравним виммера с нилом ниже он не будет равен ему, это потому что
	// значение интерфейса уже имеет конкретный тип -> уже что-то содержит
	printPlease(swimmer)
	if swimmer == nil {
		fmt.Println("swimmer is nil")
	}
	valueSwimmer := &Human{"Джордж"}
	printPlease(valueSwimmer)

	swimmer = valueSwimmer
	printPlease(swimmer)

	// пустой интрефейс - чтобв его имплементировать ненужно реализовать вообще никаих методов, мы можем положить в него вложение любого типа
	// ВАЖно понимать, если мы хотим в функции принимать любой тип, то можем использовать пустой интерфейс

	var noValueInterface EmptyInterface = NoValueSwimmer
	printPlease(noValueInterface)
}
func howToUseInterface() {
	var runner Runner // когда мы обьявляем нашу переменную интерфейсного типа Раннер мы можем в нее положить
	// как значение нашего хьюмана так и утки
	printPlease(runner)
	// И хьюман и Дак реализуют интерфейс Раннер, мы знаем, что у них есть метод РАН
	// и зхдесь, когда мы принимаем аргумент, мы не зная конкретный тип вызываем метод
	Gogi := &Human{"Gogi"}
	runner = Gogi
	printRunner(Gogi)
	typeAssertion(Gogi)

	Donald := &Duck{"Donald", "Duck"}
	runner = Donald
	printRunner(Donald)
	typeAssertion(Donald)
	//Таким образом с помощью интерфейсных значени мы можем абстрагироваться от конкретных типов
	// нам не важно -хьюман это или дак, нам важно только то, чтоу них есть метод РАН
	// таким образом ГО поддерживает полиморфизм - это когда у вас есть какая-то общая сигнатура, но у каждого типа может быть своя реализация

}

func printPlease(t any) {
	fmt.Printf("ТипТипТипчик %T, Значение: %#v\n", t, t)
}
func printRunner(runner Runner) {
	fmt.Println(runner.Run())
}

// Утверждение типа ( type assertion) - пытаемся получить значение конкретного типа

func typeAssertion(runner Runner) {
	printPlease(runner)
	if human, ok := runner.(*Human); ok { //мы создаем здесь две переменные, первое-значение конкретного типа
		//а второе (ок) - удалось ли преобразовать значение интерфесного типа в конкретный тип
		//те мы у интерфесного типа пытаемся получить конкретнвй тип
		printRunner(human)
		human.WriteCode()
	}
	if duck, ok := runner.(*Duck); ok {
		printRunner(duck)
		duck.Fly()
	}
	// ифы выше можно перезаписать более коротким способом с помощью type switch
	switch v := runner.(type) {
	case *Human:
		fmt.Println(v.Run())
	case *Duck:
		fmt.Println(v.Swim())
	default:
		printPlease(v)

	}
}

// иплементациянтерфейса - наш конкретный тип, которыей реализует все методыв которые перечисленыв в интерфейсе
// Имплементировать = реализовать методы.
// Создаем структуру хьюман
type Human struct {
	Name string
}

// Для структуры обьявляем метод РАН, который возвращает строку, те это и есть заимплементировать
func (h Human) Swim() string {
	return fmt.Sprintf("ЧЕловек %s может плавать ", h.Name)
}
