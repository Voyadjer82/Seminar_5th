package main

import "fmt"

func main() {
	var a, b float64
	var operation string

	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)
	fmt.Println("Введите операцию: +, -, * или /")
	fmt.Scan(&operation)

	switch operation {
	case "+":
		fmt.Println("Результат:", a+b)
	case "-":
		fmt.Println("Результат:", a-b)
	case "*":
		fmt.Println("Результат:", a*b)
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя")
		} else {
			fmt.Println("Результат:", a/b)
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}
