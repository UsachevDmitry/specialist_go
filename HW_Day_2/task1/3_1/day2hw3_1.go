package main

// 3.1 Пользователь вводит числа a и b (b больше a).
//     Вывести все целые числа, расположенные между ними.

import (
    "fmt"
    "math/big"
)

func main() {
    var (
        a float64
        b float64
    )
    println("Введите два числа a и b (b больше a)")
    fmt.Scan(&a, &b)
    for i := int32(a+1); i < int32(b); i++ {
        fmt.Printf("%v ", i)
    }
    if big.NewFloat(b).Cmp(big.NewFloat(float64(int32(b)))) == 1 {
        println(int32(b))
    }
}
