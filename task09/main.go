package main

import "fmt"

func main() {
	var balance, amount float64
	var action int

	fmt.Println("Введите баланс:")
	fmt.Scan(&balance)

	if balance < 0 {
		fmt.Println("Баланс не может быть отрицательным")
		return
	}

	fmt.Println("1. Посмотреть баланс")
	fmt.Println("2. Пополнить баланс")
	fmt.Println("3. Снять средства")
	fmt.Scan(&action)

	switch action {
	case 1:
		fmt.Printf("Баланс: %.2f\n", balance)
	case 2:
		fmt.Println("Введите сумму пополнения:")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Сумма должна быть больше нуля")
			return
		}
		balance += amount
		fmt.Printf("Баланс: %.2f\n", balance)
	case 3:
		fmt.Println("Введите сумму снятия:")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Сумма должна быть больше нуля")
			return
		}
		if amount > balance {
			fmt.Println("Недостаточно средств")
		} else {
			balance -= amount
			fmt.Printf("Баланс: %.2f\n", balance)
		}
	default:
		fmt.Println("Неизвестное действие")
	}
}
