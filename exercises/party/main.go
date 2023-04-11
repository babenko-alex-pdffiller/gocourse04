package main

import "fmt"

func main() {
	// Створюємо слайс масивів, який буде містити пари
	pairs := [][]string{
		{"Анна", "Ігор"},
		{"Світлана", "Василь"},
		{"Олександра", "Петро"},
		{"Юлія", "Михайло"},
	}

	// Виводимо на екран всі пари, які прийшли на вечірку
	fmt.Println("Вечірка розпочалась! Прийшли наступні гості:")
	for _, pair := range pairs {
		fmt.Printf("%s і %s\n", pair[0], pair[1])
	}

	// Пари покидають вечірку по черзі
	fmt.Printf("\n%s і %s пішли з вечірки\n", pairs[0][0], pairs[0][1])
	pairs = pairs[1:]

	fmt.Printf("%s і %s пішли з вечірки\n", pairs[0][0], pairs[0][1])
	pairs = pairs[1:]

	fmt.Printf("%s і %s пішли з вечірки\n", pairs[0][0], pairs[0][1])
	pairs = pairs[1:]

	fmt.Printf("%s і %s пішли з вечірки\n", pairs[0][0], pairs[0][1])
	pairs = pairs[1:]

	// Виводимо на екран залишившихся гостей
	fmt.Println("\nВечірка закінчилась! На залишку залишилися:")
	for _, pair := range pairs {
		fmt.Printf("%s і %s\n", pair[0], pair[1])
	}
}
