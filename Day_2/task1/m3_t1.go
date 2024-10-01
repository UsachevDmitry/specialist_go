/*
Задача №1
Вводим любое натуральное число.
Нужно посчитать сумму цифр числа, с помощью цикла for


Пример:
In: 4521
Out: 12  
*Задание 1.1: 4 + 5 + 2 + 1 = 12 - добавить к выводу сумму как выражение
*/
package main

import (
	"fmt"
)

func main() {
	var (
		number int32
		result int32 
		a string
	)
	fmt.Println("Введите натуральное число число")
	fmt.Scan(&number)

	outer:	
		for {
			a += fmt.Sprint(number%10)		
			if number > 10 {
				result += number%10
				number = number/10
				a += fmt.Sprint(" + ")
	
			} else {
				break outer
			}
		}

	fmt.Println("\nResult =",result)
	for i := len(a) - 1; i >= 0; i-- {
        fmt.Printf("%c", a[i])
    }
	fmt.Println("")
}