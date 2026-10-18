package main

import "fmt"

func main() {
	var name string = "Михаил"
	var group = "Учебная-02"
	const course string = "Бэкенд-разработка на языке Go"
	completed := 1
	var interes bool = true

	fmt.Println("Участник:", name)
	fmt.Println("Группа:", group)
	fmt.Println("Курс:", course)
	fmt.Println("Завершено работ:", completed)
	fmt.Printf("Интерес к бэкенду: %t", interes)
}
