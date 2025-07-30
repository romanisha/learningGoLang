package task_22

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Run() {
	//OsFile()
	//output()
	//input()
	//args()
	//createFile()
	//writeTo()
	fmt.Println(reverseStr("abcdefgh", 0))

}

// работа с файлами пакет OS
func OsFile() {
	newFile, err := os.Create("task_22.txt") // создаем файл, указываем имя
	if err != nil {
		log.Fatal(err)
	}
	n, err := newFile.Write([]byte("hello world")) // записали слацс байтов и строку
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)
	err = newFile.Close() // файлы необходимо закрывать6 также как и каналы
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("task_22.txt") // открывает файл ТОЛЬКО для чтения
	if err != nil {
		log.Fatal(err) // открываем существующий файл и хотим считать с него данные
	}

	buf := make([]byte, 100) // создаем буфуер и передаем его в метод РИд
	n, err = file.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf), n) // смотри, что получили

	if err = file.Close(); err != nil { //закрываем файл
		log.Fatal(err)
	}

	file, err = os.OpenFile("task_22.txt", os.O_RDWR|os.O_APPEND, 0600) // открываем снова с такими же пермишеннами как на создание, но добавляем аппенд, чтобы можно было добавить в наш файл данные
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close() // close через deffer func, чтобы не забыть закрыть его в конце проче написать через deffer
		if err != nil {
			log.Fatal()
		}
	}()

	if n, err = file.WriteString("\nhello world 2"); err != nil { // пишем новую строчку
		log.Fatal(err)
	}

	//os.ReadFile() //сокращение,чтобы не делать неск операций
	//os.WriteFile()
}

// STDOUT STDERR - файлы в операционной системе, они связаны с терминалом
// stdOut и stderr -потови вводы и вывода
func output() {
	count, err := fmt.Println("some text") // если заглянуть внутрь принтЛн, мы на самом деле вызываем Fprintln, и указываем, что мы хотим вывести данные, которые му указали stdout
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)

	count, err = fmt.Fprintln(os.Stdout, "stdout")
	if err != nil {
		log.Fatal(err)
	}

	count, err = fmt.Fprintln(os.Stderr, "stderr")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("test_file.txt") // создаем файл
	if err != nil {
		log.Fatal(err)
	}

	defer func() { // не забываем его закрыть
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	count, err = fmt.Fprintln(file, "hello\nworld")
	if err != nil {
		log.Fatal(err)
	}
}

func input() {
	var ( //  у на  есть две переменные
		text  string
		text2 string
	)

	count, err := fmt.Scan(&text, &text2) //функция для считывания даных с клавиатуры
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text, text2, count) //скан счмтал наши данные во внутрь переменных и посчитал кол-во аругментов

	count, err = fmt.Fscan(os.Stdin, &text, &text2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text, text2, count)

	file, err := os.Open("test_file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
		}
	}()

	count, err = fmt.Fscanln(file, &text) // достает построчно, возвращает кол-во переменных
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text, count)
}

// передача аргументов программе при запуске
func args() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}

// изучение интерфейсов io.ReadFrom ioWriteTo
func createFile() {
	start := time.Now()

	file, err := os.Create("test_file2.txt")
	if err != nil {
		log.Fatal(err)
	}

	var i int

	for i < 100 {
		if _, err = file.WriteString(fmt.Sprintf("%d\n", i)); err != nil { // ` записываем в файл 100 строк
			log.Fatal(err)
		}
		i++
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created file in ", time.Since(start))
}

func writeTo() { // интерфейс в пакете IO, дапустим есть файл, и мы хотим читать его в консоль в stdout
	file, err := os.Open("test_file2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	count, err := file.WriteTo(os.Stdout) // записываем файл в консоль stdout
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // после счетчика выводится кол-во байт
}

func reverseStr(s string, k int) string {
	if k == 0 || k == 1 {
		return s
	}

	//var newString []byte
	var newString string
	tf := true
	for i := 0; i < len(s); i = i + k {

		firstIndex := i
		lastIndex := i + k

		if lastIndex > len(s) {
			lastIndex = len(s)
		}

		if tf {
			newString = newString + reverse(s[firstIndex:lastIndex])
			tf = false
		} else {
			newString = newString + s[firstIndex:lastIndex]
			tf = true
		}

	}

	return newString
}

func reverse(r string) string {
	var newR string
	for i := len(r) - 1; i >= 0; i-- {

		newR = newR + string(r[i])
	}

	return newR
}
