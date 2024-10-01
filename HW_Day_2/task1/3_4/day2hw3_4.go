package main

// 3.4 В цикле получать от пользователя оценки по четырём экзаменам.
//     Вывести сумму набранных им баллов.
//     Функцию fmt.Scan() в коде используем только один раз.

import (
    "fmt"
)

func main() {
    var (
        sum int
        rating int
    )
    println("Введите 4 оценки за экзамен")
    for i := 1; i <= 4 ; i++ {
            fmt.Scan(&rating)
            sum += rating
        }
    fmt.Printf("Сумма балов = %v\n", sum)
}
