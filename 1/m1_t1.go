// Задача №1
// Программа запрашивает имя пользователя и возраст
// Нужно вывести имя и возраст за вычетом одного года
package main

import "fmt"

func main() {
	var (
		name string
		age int
	)
	fmt.Println("Введите имя пользователя и возраст")
	fmt.Scan(&name, &age)
	fmt.Println(name, age-1)
}