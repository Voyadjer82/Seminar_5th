package main

import "fmt"

func main() {
	var n, number int

	fmt.Println("Введите количество чисел:")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Количество должно быть больше нуля")
		return
	}

	sum := 0
	minimum := 0
	maximum := 0
	even := 0
	odd := 0

	fmt.Println("Введите целые числа:")
	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		sum += number

		if i == 0 {
			minimum = number
			maximum = number
		}
		if number < minimum {
			minimum = number
		}
		if number > maximum {
			maximum = number
		}

		if number%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	average := float64(sum) / float64(n)
	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее значение:", average)
	fmt.Println("Максимум:", maximum)
	fmt.Println("Минимум:", minimum)
	fmt.Println("Чётных:", even)
	fmt.Println("Нечётных:", odd)
}
