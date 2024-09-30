/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*Проверка условий по желанию
*/
package main

import (
	"fmt"
)

const cost int = 48

func main() {
	var (
		distance int32
        fuel_rate int32
	)
	fmt.Printf("Введите расстояние в км и расход в литрах на 100 км:\n")
	fmt.Scan(&distance, &fuel_rate)
	if (50 <= distance && distance <= 10000) && (5 <= fuel_rate && fuel_rate <= 25) {
		s := float64(distance) * (float64(fuel_rate)/100) * float64(cost)

		fmt.Printf("стоимость поездки в рублях = %.2f \n", s)
	} else {
		fmt.Println("Данные не корректны")
	}
}