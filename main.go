package main

import (
	"fmt"
)

func main() {
	var pass int8

	var name string
	var age uint8

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	if age <= 0 {
		pass = 0
		fmt.Println("Вы еше не роелись ")
		fmt.Print("Ваш уровень пропуска ", pass)
	}

	if age >= 18 && age < 120 {
		pass = 1
		fmt.Println("Добро пожаловать: ", name)
		fmt.Print("Вход разрешен")
		fmt.Print("Ваш уровень пропуска ", pass)
	}

	if age > 100 {
		pass = 0
		fmt.Println("Вам не может быть больше 100 лет")
		fmt.Print("Ваш уровень пропуска ", pass)
	}
}
