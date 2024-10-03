/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100) +
2) Количество (только числа) +
3) ФИО покупателя (только буквы) +
4) Контактный телефон (10 цифр) +
5) Адрес: индекс(ровно 6 цифр), город, улица, дом, квартира +

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

Реализовать конструктор и несколько методов у типа "Накладная"

Пример:
invoice = NewInvoice()

или 

order = NewOrder()

*/
package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Order struct {
	product    Product
	buyer 	   Buyer
}

type Product struct {
	name       string
	count      int64
}

type Buyer struct {
	fullName   Fullname
	phone 	   string
	address    Address
}

type Fullname struct {
	firstName  string
	lastName   string
	patronymic string
}

type Address struct {
	index 	   int64
	city 	   string
	street	   string
	house	   string
	flat  	   string
}


func main() {
	ord := NewOrder()
	fmt.Printf("Type: %T and value %v\n", ord, ord)
}

func NewOrder() *Order {
	return &Order{
		product: *NewProduct(),
		buyer: *NewBuyer(), 	   
    }
}

func NewProduct() *Product {
	return &Product{
		name: inputStr("Введите название товара"),
		count: inputInt("Введите количество товара"),
	}
}

func NewBuyer() *Buyer {
	return &Buyer{
		fullName: *NewFullname(),
		phone: inputStr("Введите телефон"),
		address: *NewAddress(),
	}
}

func NewFullname() *Fullname {
	return &Fullname{
		firstName:  inputStr("Введите Имя"),
		lastName:   inputStr("Введите Фамилию"),
		patronymic: inputStr("Введите Отчество"),
	}
}

func NewAddress() *Address{
	return &Address{
		index:  inputInt("Введите Индекс"),
		city:   inputStr("Введите Город"),
		street: inputStr("Введите Улицу"),
		house:	inputStr("Введите дом"),
		flat:  	inputStr("Введите квартиру"),
	}
}


func inputStr(text string) string {
	fmt.Println(text)
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	return str
}

func inputInt(text string) int64 {
	fmt.Println(text)
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	count, _ := strconv.ParseInt(strings.TrimSuffix(str, "\n"), 10, 64)
	return count
}