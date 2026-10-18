package main

import "fmt"

func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

func countStatus(statuses []string, target string) int {
	var count int
	for _, w := range statuses {
		if w == target {
			count++
		}
	}
	return count
}

func findStatus(statuses []string, target string) int {
	for i, w := range statuses {
		if w == target {
			return i
		}
	}
	return -1
}

func completionPercent(statuses []string) float64 {
	doneElements := float64(countStatus(statuses, "done"))
	lenSlice := float64(len(statuses))
	if lenSlice == 0 {
		return 0
	}
	return 100 * doneElements / lenSlice
}
func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Println("Задач всего:", len(statuses))

	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Println("Завершено:", countStatus(statuses, "done"))

	counts := make(map[string]int)

	for _, w := range statuses {
		counts[w]++
	}
	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d\n", counts["new"], counts["in_progress"], counts["done"])

	find := findStatus(statuses, "done")
	if find == -1 {
		fmt.Println("Элемент не найден")
	} else {
		fmt.Printf("Совпадение на %d индексе\n", find)
	}

	fmt.Printf("Процент завершенных задач: %.2f", completionPercent(statuses))
}
