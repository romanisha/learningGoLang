package task_22

import (
	"fmt"
	"log"
	"os"
)

func Run() {
	//OsFile()
	output()
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
