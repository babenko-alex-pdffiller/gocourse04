package main

import "fmt"

func main() {
	// Створюємо слайс, що містить перший вагон потягу
	train := []string{"Вагон 1"}

	// Додаємо по три вагона на кожній станції
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			wagon := fmt.Sprintf("Вагон %d-%d", i+1, j)
			train = append(train, wagon)
		}
	}

	// Виводимо потяг
	fmt.Println(train)
}
