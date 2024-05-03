package main

import "fmt"

func main() {
	for {

		var action string
		var a float64
		var b float64
		fmt.Println("Введите задачу, например 2 * 2:")
		fmt.Scan(&a, &action, &b)

		switch {
		case action == "+":
			fmt.Println("Сумма чисел" + " " + fmt.Sprint(a+b))
		case action == "-":
			fmt.Println("Разность чисел" + " " + fmt.Sprint(a-b))
		case action == "*":
			fmt.Println("Произведение чисел" + " " + fmt.Sprint(a*b))
		case action == "/":
			fmt.Println("Частное чисел" + " " + fmt.Sprint(a/b))
		default:
			{
				fmt.Println("Неверное действие")
			}

		}
	}
}
