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
		sum += i
	}

	average := float64(sum) / float64(n)
	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее значение:", average)
}
