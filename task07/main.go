package main

import "fmt"

func main() {
	var n int

	fmt.Println("Введите положительное целое число N:")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("N должно быть больше нуля")
		return
	}

	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}

	fmt.Println("Сумма подходящих чисел:", sum)
}
