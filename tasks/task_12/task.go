package task_12

import "fmt"

func Run() {
	explanation()
	Usecase()
}

// встраивание не равно наследование
// встраивание работает по композиции, потому что в го нет наследования вообще
type Person struct {
	name string
	age  int
}
type WoodBuilder struct { // мы хотим в него встроить наш тип персон
	Person // те мы встраиваем сюда персон
	//name   string // добавляем имя еще и самому вудбилдеру
	//WorkExperience
}
type BrickBuilder struct {
	Person
}

type WorkExperience struct {
	name string
	age  int
}
type Building struct {
	name string
	Builder
}

type Builder interface {
	Build()
}

func (wb WoodBuilder) Build() {
	fmt.Println("Строю дом из дерева")
}

func (bk BrickBuilder) Build() {
	fmt.Println("строю дом из кирпичиков говна")
}
func (p Person) PrintName() {
	fmt.Println(p.name)
}

func (wb WoodBuilder) PrintName() {
	fmt.Println(wb.name)

}

func explanation() {
	builder := WoodBuilder{Person{"Dava", 37}}
	//builder := WoodBuilder{Person{name: "Dava", age: 37}, "Vasya"}
	//builder := WoodBuilder{Person{name: "Dava", age: 37},
	//	"Vasya",
	//	WorkExperience{name: "Nastya", age: 25}} //colliding
	printPlease(builder)

	//shorthands. Получение свойств, вызов методов
	// чтобы вывести возраст:
	fmt.Println(builder.Person.age)
	//fmt.Println(builder.age) // синтаксический сахар или  шортхэнд
	//в случае с Настей го имеет age в двух составляющих одного уровня - как у Person, так и у WorkExperience.
	//fmt.Println(builder.WorkExperience.name)

	//shadowing
	fmt.Println(builder.name) // в данном случае го найдем имя вася- потому что он не проваливается дальше во внутрь из-за 30 строчки

	// вызов методов
	builder.PrintName()
}

func Usecase() {
	woodenBuilding := Building{
		Builder: WoodBuilder{Person{
			name: "Dava",
			age:  37,
		}},
		name: "Деревянная изба",
	}
	woodenBuilding.Build()

	brickBuilding := Building{
		Builder: BrickBuilder{
			Person{
				name: "Petuh",
				age:  22,
			},
		},
		name: "Кирпичиковый дом",
	}
	brickBuilding.Build()
}

func printPlease(t any) {
	fmt.Printf("ТИП: %T, Значение %#v\n", t, t)
}
