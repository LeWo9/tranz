package main

import (
	"fmt"
)

func main() {
	// Запрос ввода строки от пользователя
	fmt.Print("Введите строку: ")
	var input string
	fmt.Scanln(&input)

	// Создание слайса, содержащего каждый символ строки
	charSlice := make([]rune, 0)
	for _, char := range input {
		charSlice = append(charSlice, char)
	}

	// Вывод слайса на экран
	fmt.Printf("Слайс символов: %v\n", charSlice)
}
