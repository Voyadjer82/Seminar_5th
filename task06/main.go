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

	positive := 0
	negative := 0
	zeros := 0

	fmt.Println("Введите числа:")
	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if number > 0 {
			positive++
		} else if number < 0 {
			negative++
		} else {
			zeros++
		}
	}

	fmt.Println("Положительных:", positive)
	fmt.Println("Отрицательных:", negative)
	fmt.Println("Нулей:", zeros)
}
