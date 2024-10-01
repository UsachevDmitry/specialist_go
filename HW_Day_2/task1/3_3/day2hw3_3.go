package main

//3.3 Пользователь вводит число. Вывести таблицу умножения на это число (от 1 до 10)

import (
    "fmt"
)

func main() {
    var (
        number int
    )
    println("Введите число")
    fmt.Scan(&number)
    for i := 1; i <= 10 ; i++ {
            fmt.Printf("%v * %v = %v\n", number, i, number * i)
        }
}
