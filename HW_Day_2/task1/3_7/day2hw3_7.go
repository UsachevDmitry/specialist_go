package main

// 3.7 Вводим строку без знаков препинания(5 слов).
//     Найти самое длинное слово в строке и вывести сколько в нем букв.

// Пример:
// In: Скажи как дела в учебе 
// Out: учебе, 5

// In: Закрепляем циклы в языке Golang
// Out: Закрепляем, 10
// */


import (
    "fmt"
    "unicode/utf8"
)

func main() {
	var (
		str string = "Скажи как дела в учебе"
        word string = ""
        max_word string
        max int
        i int
        max_word_count int = 1
	)
    str = str + "!" // костыль чтобы обработать последний символ
    for _ , char := range str {    
        i++    
        if char != ' ' && i != utf8.RuneCountInString(str) {   
            word = word + string(char)
        } else {    
            if utf8.RuneCountInString(word) > max {
                max = utf8.RuneCountInString(word)
                max_word = word
            } else if utf8.RuneCountInString(word) == max {
                max_word_count ++ 
            }
        word = ""
        }            
    }
	fmt.Printf("Самое длинное первое слово = '%v'\nЧисло символов в этом слове = %v\nЧисло слов максимальной длинны = %v\n", max_word, max, max_word_count)
}
