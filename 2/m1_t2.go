// Задача №2
// Вход: трехзначное число. 
// Найти первую, вторую и последнюю цифры заданного трехзначного числа.
package main

import "fmt"

func main() {
	var (
		c int
	)
	fmt.Println("Введите 3х значное число")
	fmt.Scan(&c)
	first := c/100
	second := c%100/10
	fmt.Printf("Наше число = %v , первая цифра = %v, вторая цифра = %v\n",c, first, second)
}