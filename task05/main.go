package main

import "fmt"

func main() {
	var n int
	var number float64

	fmt.Println("Введите количество чисел:")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Количество должно быть больше нуля")
		return
	}

	fmt.Println("Введите числа:")
	fmt.Scan(&number)
	minimum := number
	maximum := number

	for i := 1; i < n; i++ {
		fmt.Scan(&number)
		if number < minimum {
			minimum = number
		}
		if number > maximum {
			maximum = number
		}
	}

	fmt.Println("Максимум:", maximum)
	fmt.Println("Минимум:", minimum)
}
