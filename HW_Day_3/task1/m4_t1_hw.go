/* 
	Задача №1
	Написать функцию, которая расшифрует строку. 
	code = "220411112603141304"
	Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
	Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
	Вход: строка из цифр. Выход: Текст.
	Проверка работы функции выполняется через другую строку.

	Задача №1.1
	Реализовать и функцию зашифровки

	codeToString(code) -> "???????' 

	stringToCode("hello") -> "??????"
*/

package main

import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	var (

		abc string = "abcdefghijklmnopqrstuvwxyz "
	)
	mapCode := make(map[string]string)
	// Создаем map c кодами
	for i, char := range abc {
		mapCode[fmt.Sprintf("%02d", i)] = fmt.Sprintf("%v", string(char))
	}
	fmt.Println("Наша мапа шифрования", mapCode)

	fmt.Println("Введите строку для шифровки:")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Ваша зашифрованная строка:", stringToCode(str, mapCode))

	fmt.Println("Введите строку для расшифровки:")
	str2, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Ваша расшифрованная строка:", codeToString(str2, mapCode))
}

func stringToCode(str string, mapCode map[string]string) string {
	var answer string
	for _, char := range str {
		for key := range mapCode {
			if string(char) == mapCode[key] {
				answer += key
			}
		}
	}
	return answer
}

func codeToString(str string, mapCode map[string]string) string {
	var answer string
	for i := range str {
		if i % 2 != 0 {
			key_symbol := string(str[i-1]) + string(str[i])
			for key, value := range mapCode {
				if string(key_symbol) == key {
					answer += value
				}
			}
		}
	}
	return answer
}