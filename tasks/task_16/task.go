package task_16

import (
	"fmt"
	"time"
)

func Run() {
	//nilChannel()
	//unbufferedChannel()
	//bufferedChannel()
	forRange()

}

//каналы - инструммент коммуникации, который помогает обмениваться данными между горутинами
// канал -тип данных в го, вызывается как channel, внутри канала находится mutex
//type chan struct {
//	mx sync.mutex
//	buffer [T]  - может быть буфер, а может не быть. смотря буферезированный он или нет
//	readers []Goroutines
//	writers []Goroutines
//}

func nilChannel() {
	// creating nil channel
	var nilChannel chan int
	getLength(nilChannel) // у канала есть длина(кол-во элементов, которые находятся в канале) и капасити(размер буфера канала)

	//записывать  значение из NIL канала - невозможно
	nilChannel <- 1 // помещаем туда значение 1, это значение некуда поместить, в нем нет ничего ни буфера, ни горутин

	//чтение из NIL канала -невозможно
	<-nilChannel

	//close NIL channel
	close(nilChannel) // итого ниловый канал мы почти не будем использовать, в него нельзя записать, прочитать, закрыть его нельзя
}

func unbufferedChannel() {

	//creating unbuffered channel with make
	unbufferedChannel := make(chan int) // двунаправленный канал, те и читает и записывает
	getLength(unbufferedChannel)

	//небуфферезированный канал - это канал, буффер которого равен 0, те он не используется и значения в нем не хранятся
	//unbufferedChannel <- 1 // записывем значение в канал

	//<-unbufferedChannel // чтение из канала

	//blocks on reading then write
	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		//<- chanForWriting
		unbufferedChannel <- 1
	}(unbufferedChannel)

	value := <-unbufferedChannel
	fmt.Println(value)

	//blocks on writing then reading

	go func(chanForReading <-chan int) {
		time.Sleep(time.Second)
		value := <-unbufferedChannel
		fmt.Println(value)
	}(unbufferedChannel)

	unbufferedChannel <- 2

	//close channel
	go func() {
		time.Sleep(time.Millisecond * 500)
		close(unbufferedChannel)
	}()
	go func() {
		time.Sleep(time.Second)
		fmt.Println(<-unbufferedChannel)
	}()
	unbufferedChannel <- 3 // поймаем панику тк канал закроется

	// close of closed panics
	close(unbufferedChannel) // закрытие закрытого канал, лучшая праквтика закрывать канал там, где в него пишутся данные
	close(unbufferedChannel)

	//направленность канала, допустим нам нужен канал ТОЛЬКО для записи
	//unbufferedChannel1 := make(chan<- int) // означает, что мы можем только записывать значения в канал
	//unbufferedChannel2 := make(<-chan int) // означает, что мы можем только читать значения в канале
}

func bufferedChannel() {
	//creating buffered channel
	bufferedChannel := make(chan int, 2) // второй аргумент - размер буфера
	getLength(bufferedChannel)

	//не блокирует запись, пока буфер не полный
	bufferedChannel <- 1
	bufferedChannel <- 2
	getLength(bufferedChannel)

	//блокирует запись, когда буфер полный
	bufferedChannel <- 3 // уже заблокируемся

	// пытаемся достать значения из канала, и у нас это получаетсч, тк есть значения, которые можн извлечь
	// в небуферезированном канале мы сначала записывали, потом доставали
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	getLength(bufferedChannel)

	//блокирует чтение, буффер пустой, нет записей, тк мы уже достали все значение из буфера
	fmt.Println(<-bufferedChannel)
}

func forRange() {
	bufferedChannel := make(chan int, 3)

	numbers := []int{5, 6, 7, 8}

	//show all elements with FOR
	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel) // закрываем канал после завершения записи
	}()

	for {
		v, ok := <-bufferedChannel // ок проверяет открыт или закрыт канал, также как у мап тру иди фолс
		fmt.Println(v, ok)
		if !ok { // если канал закрыт -прекращаем работу
			break
		}
	}

	//создаем буфферезированный канал, тк в первом кейсе мы его закрыли
	bufferedChannel = make(chan int, 3)

	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel)
	}()

	for v := range bufferedChannel { // из-за флренджа мы знаем, что канал закрылся и никаких доп проверок не нужно
		fmt.Println("buffered", v)
	}

	// пример for range  с небуфферизмрованным каналом
	//пропускная способность всего одно значение в отличие от пред примера
	unbufferedChannel := make(chan int)
	go func() {
		for _, num := range numbers {
			unbufferedChannel <- num
		}
		close(unbufferedChannel)
	}()

	for val := range unbufferedChannel {
		fmt.Println("unbuffered", val)
	}
}

func getLength(channel chan int) {
	fmt.Printf("Length: %d Capacity: %d\n\n", len(channel), cap(channel))
}
