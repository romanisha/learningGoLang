package task_15

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Run() {
	//HowUseGoroutine()
	//WhatIsDeferredFunc()
	//deferValues()
	//makePanic()
	//withoutWait()
	//withWait()
	//wrongAdd()
	//writeWithoutConcurrency()
	//writeWithoutMutex()
	//writeWithMutex()
	readWithMutex()
	readWithRWMutex()

}

func HowUseGoroutine() {

	// на копьютере может выполняться столько горутин, сколько в вашем процессоре логияеских ядер
	fmt.Println(runtime.NumCPU()) // вычисляем сколько у меня логических ядер, те может одновременно параллельно выполнятся до 8 горутин

	// мы можем повлиять на количество горутин, которые у нас выполняются одновременно
	runtime.GOMAXPROCS(1) //ставим кол-во горутин, которое могут выполнятся на компютере одновременно

	// горутины могут выполняться парралельно и конкурентно
	//для правильной работы с горутинами надо синхронизировать
	//горутина может иметь два сострояния: выполняется или заблокирована
	go showNumbers(100)

	// мы можем вручную говорить планировщику го когда хотим переключится на дргую горутину
	//runtime.Gosched() //тут мы переключились на нашу горутину и горутина завершилась

	// Перкключением между горутинами занимается планировщик го и он сам принимает решение на основее кода который у нас есть
	time.Sleep(time.Second)

	fmt.Println("exit")
}

func WhatIsDeferredFunc() {
	//Вызывается словом deffer, эти функции не выполняются сразу как только мы вызвали, они складываются в stack defer функции
	// и начинают выполнятся после того, как выполняется конструкция ретерн, при чем они выполняются в обратном порядке
	defer fmt.Println(1)
	defer fmt.Println(2)

	// возможность defer функции изменять значение возвращаемых аргументов ( именно тех, которые были именованы)
	fmt.Println(sum(2, 3))

	fmt.Println("exit")
}

func deferValues() {
	//Вычисление значений деферед фанкции
	//2 варианта вызова
	for i := 0; i < 10; i++ {
		defer fmt.Println("First", i)
	}
	for i := 0; i < 10; i++ { //С версии Go 1.22  на каждой итерации цикла создается новая переменная, а не изменяется значения одной и той же переменной i.
		defer func() { //Поэтому defer забирает каждый раз обращается к разным ячейкам памяти и выводит разный результат.
			fmt.Println("Second", i) //Раньше же все defer'ы обращались к одной и той же ячейки памяти, в которой после окончания цикла оставалась переменная со значением 10
		}()
	}
}

func makePanic() {
	defer func() {
		panicValue := recover() //можно отловить панику с помощью фукнкции recover
		fmt.Println(panicValue) //результат функции рековер будет значение паники, если она произошла и НИЛ, если нет
	}()

	panic("some panic") //  поника может быть, если вызвать элементы за рамками массива или слайся
	// также паника может быть вызвана nil pointer exeption , те у нила пытаемся вызвать функцию у нила
	//когда выбрасывается паника наша текущая функция завершается и горутина старается завершиться тоже
	fmt.Println("Unreachable code") //после того, как случается паника наш код перестает выполнятся
}

func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1) //функция в цикле создает 10 горутин, которые просто распечатают наш счетчик
	}
	fmt.Println("exit") // мы ожидаем, что увидим вывод из 10 числе и эксит
	// но выводятся не всегда все 10 чисел, тк горутины не успевают запуститься и мы их не ждем в нашей основно горутине
}

func withWait() {
	//wait group - механизм ожидания завершения группы задач
	var wg sync.WaitGroup //создаем переменную у которой тип вейт групп из пакета синк
	//у вейтгруппа есть несколько методов - add done wait
	// если известно заранее кол-во задач, то можно добавить в вецт группу заранее wg.Add(10) вместо 94 строчки

	for i := 0; i < 10; i++ {
		wg.Add(1) // метод принимает кол-во задач кторое нужно добавить в нашу вейтгруппу те добавляем 10 задач

		go func(i int) { // на каждую итерации мы добавляем новую горутину, передаем в нее счетчик цикла
			fmt.Println(i) // наша горутина выводит в консоль счетчик
			wg.Done()      // важно обязательно вызывать метод done, он сообщаяет о том, что наша таска завершилась
		}(i)
	}
	wg.Wait() // наша основная горутина блокируется и ждет пока в вейт группе не останется невыполненных задач
	fmt.Println("exit")
}

func wrongAdd() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func(i int) {

			wg.Add(1) //добавляем задачу в нашу вейтгруппу изнутри горутины, а не снаружи
			// горутина не успевает внести себя в список ожидания, она просто пропускается и выполняется
			defer wg.Done() //поэтому лучше добавлять задачи в вейт группу заранее как в предвдущем примере

			fmt.Println(i + 1)
		}(i)
	}

	wg.Wait()
	fmt.Println("exit")
}

func writeWithoutConcurrency() {
	start := time.Now()
	var counter int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond)
		counter++
	}

	fmt.Println(counter)
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds()) //мы выполнили 1000 итераций за 1,2 сек
}

func writeWithoutMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup // почему ругается при создании через литерал wg := sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ { // мы ожидаем, что счетчик будет 1000, но какие-то обнавления счетчика мв потеряли
		// это происходит из-за data race(обращение к одним и тем же данным из разных программ/тредов/горутин
		//где хотя бы одно из этих обращений - это запись
		// получается каждая из горутин устанавливает одно и то же значение, они выполнили одну и туже работу
		go func() { // мы запускаем код в 1000 горутинах
			defer wg.Done() // закрвваем нашу таску используюя дефер
			time.Sleep(time.Nanosecond)
			counter++ //
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())
}

// для решения проблем с data race нужен mutex
func writeWithMutex() {
	start := time.Now()
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex // если мы используем мьютекс мы можем получить блокировку на какой-то код
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock() // позволяет заблокировать участок кода, чтобы несколько горутин одновременно не записывали/ не читали из каунтера
			counter++ //с блокировкой мы уверены,что с этим участком кода работает только одна горутина, те все остальные 999 горутин ждут
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())
}

func readWithMutex() { // представим, что есть 50 операций чтения и 50 операций записи
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)
	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			_ = counter
			mu.Unlock()
		}()
		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())
}

func readWithRWMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.RWMutex
	)
	wg.Add(100)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.RLock() // используется для чтения, получает НЕ эксклзивнкю блокировку, поэтому читаем быстрее
			time.Sleep(time.Nanosecond)
			_ = counter
			mu.RUnlock()
		}()

		go func() {
			defer wg.Done()
			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())
}

func showNumbers(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func sum(x, y int) (sum int) { // используем именнованное возвращаемое значение sum
	defer func() { // вызов deferred функции в котором мы берем sum и умножаем на 2
		sum *= 2
	}()
	sum = x + y
	return // после ретерна вызывается deferred функция
}
