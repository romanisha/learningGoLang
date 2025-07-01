package task_18

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Run() {
	//learningContext()
	workerPool()
}

// context -входит в стандартную библиотеку, контекс - это такой обьект, которыц служит для двух целей
// 1ая -хранить значение, что бы дальше передавать по цепочке во вложенные функции или горутины
//2ая -возможность сообщать о завершении, те изнутри самого контекста мы можем узнать его состояние (завершен контекст или нет)

//смысл контекста заключается в том, что мы изначально создаем где-то контекс наверху, примерно в корне нашего приложения
//часто происходит к примеру приходит какой-то http запрос и там у нас формируется контекс и этот контекс дальше идет по всем функциям
//также  мы можем в рамках запроса подложить какие-нибудь значения, допустим инфу о пользователе и тд
//и можем сообщать о завершении

func learningContext() {

	//creating context
	//можно создать два вида контекста - context.Background и context.TODO
	// разница: считается, что вариант toDo пригоден только в тестах
	// Context.Background создает наш корневой контекст, далее мы можем создавать новые контексты на основе родительского
	ctx := context.Background()
	fmt.Println(ctx)

	toDo := context.TODO()
	fmt.Println(toDo)

	withValue := context.WithValue(ctx, "name", "vasya")
	fmt.Println(withValue.Value("name"))

	//  контекс как инструмент, которыц сожет сообщить о завершении какой-то задачи, отмена контекста

	withCancel, cancel := context.WithCancel(ctx) // создаем контекст, который умеет завершаться, если мы создадим контекст, используя функцию WithCancel
	//здесь подразумевается, что мы создаем контекст, который мы можем отменить, если мы исполььзуем WithCancel, то мы получаем на выходе контекст
	//и мы получаем 1) withCancel - новый контекст 2)cancel - функция, которая умеет завершать этот контекст
	fmt.Println(withCancel.Err()) // вызыввем метод еррор, хотим посмотреть есть ли в контексте ошибка
	cancel()                      // вызываем, чтобы завершить контекст
	fmt.Println(withCancel.Err()) // видим ошибку о том, что контекст завершен

	// контекст с установленным дед лайном
	withDeadLine, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3)) //в качестве дедлайна принимаем duration те временную ветку
	// и мы передаем время, когда контекст должен завершиться, те получаю текущее время и прибавляю  нему 3 сек (через 3 сек контекст должен завершиться)
	defer cancel() //  как только функция завершится наш контекс отменится
	fmt.Println(withDeadLine.Deadline())
	fmt.Println(withDeadLine.Err())
	fmt.Println(<-withDeadLine.Done()) // метод done возращает канал,но в этом канале значения нет, поэтому выводится пустая структура ЧЕРЕЗ 3 сек

	withTimeOut, cancel := context.WithTimeout(ctx, time.Second*3) //похоже не дедлайн
	defer cancel()                                                 //если уберем закрвтие контекста по выходу из функции, иде будет ругаться
	fmt.Println(<-withTimeOut.Done())                              // получаем канал - пустая структура
}

func workerPool() {
	//  workerPool паттерн использования конкурентности в го

	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст, который ожно отменять, на основе род. контекста бэкраунд
	defer cancel()

	wg := sync.WaitGroup{}
	numbersToProcess, processedNumbers := make(chan int, 5), make(chan int, 5) // содаем буферезованные каналы

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			//if i == 500 { // как только обработали 500 берем и завершаем наш контекст
			//	cancel()
			//}
			numbersToProcess <- i //кладем цифры в первый канал
		}
		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()               //ждем пока воркеры завершат работу с помощью вейт группы
		close(processedNumbers) // закрываем канал для корректной работы
	}()

	var counter int // далее создаем перемунную , которое считает, сколько значений обработали
	for resultValue := range processedNumbers {
		counter++
		fmt.Println(resultValue)
	}

}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select { //  у селекта есть два пути
		case <-ctx.Done(): //1ый дождаться завершения контекста, эта операция блокирующая
			return
		case value, ok := <-toProcess: //2ой неблок, он предпочтительнее, берем канал из которого читаем цифры
			if !ok { // если канал закрылся, выходим, чтобы не зависнуть
				return
			}
			time.Sleep(time.Millisecond) // после получения значения из канала наш воркер спит одну милсекунды
			processed <- value * value   // и получаемое число возводим в квадрат и переложили в др канал

		}
	}
}
