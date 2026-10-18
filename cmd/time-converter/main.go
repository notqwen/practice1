package main

import "fmt"

func main() {
	var totalMinutes int
	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&totalMinutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}
	const minutesInHour int = 60
	var totalHours = totalMinutes / minutesInHour
	fmt.Printf("Полных часов: %d\n", totalHours)
	fmt.Printf("Оставшихся минут %d", totalMinutes%minutesInHour)
}
