/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1 9 2
Выход: 1 2 9

Про sort мы пока ничего не знаем!
Это задача на условный оператор
*/
package main

import (
	"fmt"
)

func main() {
	var (
		number1 int32
		number2 int32
		number3 int32
		first int32
		midle int32
		last int32
	)

	fmt.Println("Введите три числа")
	fmt.Scan(&number1, &number2, &number3)
	if number1 > number2 && number1 > number3 {
		first = number1
		if number2 > number3 {
			midle = number2
			last = number3
		} else {
			midle = number3
			last = number2
		}
	} else if number2 > number1 && number2 > number3 {
		first = number2
		if number1 > number3 {
			midle = number1
			last = number3
		} else {
			midle = number3
			last = number1
		} 
	} else if number3 > number1 && number3 > number2 {
		first = number3
		if number1 > number2 {
			midle = number1
			last = number2
		} else {
			midle = number2
			last = number1
		} 
	}
	fmt.Printf("Три числа по возрастанию %v %v %v \n", last, midle, first)
}