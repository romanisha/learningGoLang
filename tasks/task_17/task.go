package task_17

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	//baseSelect()
	gracefulShutDown()
}

func baseSelect() {
	bufferedChan := make(chan string, 3) //1
	bufferedChan <- "first"

	select { //свитч кейс для работв с каналами,различает три типа операции block, unblock, default
	// в первую очередь селект анализирует все наши case и ищет Неблокирующие операции! если их несколько он возьмет их в рандомном порядке
	case str := <-bufferedChan: //это непблокирующая операция, тк в буфере значение есть и мы хотим прочитать
		fmt.Println("read", str)
	case bufferedChan <- "second": //это блокирующая операция, тк буфер заполнен
		fmt.Println("write", <-bufferedChan, <-bufferedChan)
	}

	unbufChan := make(chan int)

	go func() { //записываем значение через секунду
		time.Sleep(time.Second)
		unbufChan <- 1
	}()
	select {
	case bufferedChan <- "third": //запись в буферезированныц канал
		fmt.Println("unblocking writing") // приоритет отдан ей, она неблокирующая, но если этот кейс убрать, то будет деыолт, тк выполняется мнгновенно
	case val := <-unbufChan: //блокируемое чтение
		fmt.Println("blocking reading", val)
	case <-time.After(time.Millisecond * 1500): // при значени 1,5 секунды таймаут наступить не успевает, и происходит предыдущий кейс
		fmt.Println("timeout") // по истечении таймера у нас будет отпрвлено значение в канал
	default:
		fmt.Println("default") // дефолт выполняется за секунду, приоритет отдан ему. но если отключить, то приоритет пойдет ко второму кейсу, тк он выполняется быстрее таймаута
	}

	resultChan := make(chan int)
	timer := time.After(time.Second) // если запускаем наш селект в цикле нужно таймер выносить наружу

	go func() {
		defer close(resultChan)

		for i := 1; i <= 1000; i++ { // предположи у нас 100 операций и на каждую требуется наносекунда (54 строка)
			// мы не хотим ждеть пока выполнятся все операции, а просто хотим выполнить столько,сколько успеем за опред кол-во времени
			select {
			case <-timer: // когда время кончается мы выходим
				fmt.Println("timeout")
				return
			default: // пока время есть мы вопалняем дефолтную операцию
				time.Sleep(time.Nanosecond)
				resultChan <- i // и записываем
			}
		}
	}()
	for v := range resultChan {
		fmt.Println(v)
	}
}

func gracefulShutDown() { //возможность завершить программу дав ей время на закрытие соединений и выполнение операций
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // либо  мы посылаем ей сигнал, она завершается после сигнала

	timer := time.After(10 * time.Second) // программа либо сама завершается через 10 секунд

	for {
		select {
		case <-timer:
			fmt.Println("timeout")
			return
		case sig := <-sigChan:
			fmt.Println("stopped by signal", sig)
			return
		}
	}
}
