package main

import "fmt"

func main() {
	var num1 int = 20
	var num2 int = 25
	var num3 int = 35

	var totalMinutes float64 = float64(num1 + num2 + num3)
	fmt.Printf("Сумма элементов %.f\n", totalMinutes)
	var countMinutes float64 = 3
	var midMinutes float64 = totalMinutes / countMinutes
	fmt.Printf("Среднее значение: %.2f", midMinutes)
}
