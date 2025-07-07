package task_19

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"sync/atomic"
	"time"
)

func Run() {
	//chanAsPromise()
	//chanWithMutex()
	//withOutErrGroup()
	//errGroup()
	//AddMutex()
	//addAtomic()
	//StoreLoadSwap()
	//compareAndSwap()
	atomicValue()
}

func makeRequest(num int) <-chan string {
	responseChan := make(chan string)

	go func() {
		time.Sleep(time.Second) // имитация http запроса
		responseChan <- fmt.Sprintf("responce number %d", num)

	}()
	return responseChan
}

// использование канала как промиса
func chanAsPromise() {
	firstResponceChan := makeRequest(1)
	secondResponceChan := makeRequest(2)
	// do smth else
	fmt.Println("not blocking")

	fmt.Println(<-firstResponceChan, <-secondResponceChan)
}

// использование канала как Мьютекса
func chanWithMutex() {
	var counter int
	mutexChan := make(chan struct{}, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mutexChan <- struct{}{}
			counter++
			<-mutexChan
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func withOutErrGroup() {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() { // первая горутина не выполнится, тк во время ее сна во второй горутине уже  произойдет кансел и контекс завешен
		time.Sleep(time.Second)
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("first started")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("second started")
			err = fmt.Errorf("any error")
			cancel()
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("third started")
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	fmt.Println(err)
}

func errGroup() {

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error { //аналогично предыдущей функции данная горутина не успевает выполнится
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			return nil
		default:
			fmt.Println("first started")
			time.Sleep(time.Second)
			return nil
		}
	})
	g.Go(func() error {
		fmt.Println("second started")
		return fmt.Errorf("unexpected error in request 2")
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("third started")
			time.Sleep(time.Second)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

// изучаем пакет синк/атомик
// методв данного пакета различается в зависимости от типа ( включает в себя обработку ТОЛЬКО интовых значений)
// всего 5 методов : Add, Load, Store, Swap, compareAndSwap
// пакет синк занимался синхронзацией между горутинами (waitGroup) и у нас был Mutex ( чтобы не случались data race)n
// атомик работает быстрее, потому что в пакете используются операции на уровне процессора. Но минус пакета - можете использовать только одну операцию
func AddMutex() {
	start := time.Now()
	var (
		counter int64
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("With mutex:", time.Now().Sub(start).Seconds())
}

func addAtomic() {
	start := time.Now()

	var (
		counter int64
		wg      sync.WaitGroup
	)
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1) // дельта - значение, на которое нужно изменить наше значение. передаем указатель на счетчик ( каунтер)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("With atomic:", time.Now().Sub(start).Seconds())
}

// изучаем как работает Store Load Swap
func StoreLoadSwap() {
	var counter int64

	fmt.Println(atomic.LoadInt64(&counter)) //Load получает значение атомарно

	atomic.StoreInt64(&counter, 5)          // Store кладет значения атомарно
	fmt.Println(atomic.LoadInt64(&counter)) // смотрим что выведется, получаем 5

	fmt.Println(atomic.SwapInt64(&counter, 10)) // выдает сначала старое значение, потом новое (5, 10)
	fmt.Println(atomic.LoadInt64(&counter))
}

func compareAndSwap() {
	var (
		counter int64
		wg      sync.WaitGroup
	)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()

			if !atomic.CompareAndSwapInt64(&counter, 0, 1) { // нужно изменить значение с 0 на 1.
				//1ый аргумент -указатель на счетчик, 2ым -старое значение,3им- новое значение. Возвращается bool, которая называется swap
				// он возвращает тру, только если значение получилось поменять
				return
			}
			fmt.Println("Swapped goroutine number is: ", i)
		}(i)
	}
	wg.Wait()
	fmt.Println(counter)
}

func atomicValue() {
	var (
		value atomic.Value //  тип atomic.Value имеет пустую структуру, значит можно помещать не только интовые значения
	)
	value.Store("hello world")
	fmt.Println(value.Load())

	value.Swap("privet")
	fmt.Println(value.Load())
	fmt.Println(value.CompareAndSwap("privet", "lol")) // хотим заменить старое значение на новое и сравнить
	fmt.Println(value.Load())
}
