package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var availableCurrencies = []string{"USD", "EUR", "RUB"}

func isValidCurrency(currency string) bool {
	for _, c := range availableCurrencies {
		if c == currency {
			return true
		}
	}
	return false
}

func inputCurrency(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s (доступно: %s): ", prompt, strings.Join(availableCurrencies, ", "))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToUpper(input))

		if isValidCurrency(input) {
			return input
		}

		fmt.Println("❌ Ошибка: такой валюты нет. Попробуйте снова.")
	}
}

func inputNumber(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		number, err := strconv.ParseFloat(input, 64)
		if err == nil && number > 0 {
			return number
		}

		fmt.Println("❌ Ошибка: введите корректное положительное число.")
	}
}

func calculate(amount float64, from string, to string) float64 {

	rates := map[string]float64{
		"USD": 90.0,
		"EUR": 100.0,
		"RUB": 1.0,
	}

	inRub := amount * rates[from]
	result := inRub / rates[to]
	return result
}

func main() {
	fmt.Println("💰 Калькулятор валют")
	fmt.Println("\nШаг 1: Выберите исходную валюту")
	from := inputCurrency("Исходная валюта")

	fmt.Println("\nШаг 2: Введите сумму")
	amount := inputNumber("Сумма: ")

	fmt.Println("\nШаг 3: Выберите целевую валюту")
	to := inputCurrency("Целевая валюта")

	var result float64
	switch {
	case from == to:
		result = amount
	case from == "USD" && to == "EUR":
		result = amount * 90.0 / 100.0
	case from == "USD" && to == "RUB":
		result = amount * 90.0
	case from == "EUR" && to == "USD":
		result = amount * 100.0 / 90.0
	case from == "EUR" && to == "RUB":
		result = amount * 100.0
	case from == "RUB" && to == "USD":
		result = amount / 90.0
	case from == "RUB" && to == "EUR":
		result = amount / 100.0
	default:
		result = calculate(amount, from, to) // резервный вариант
	}

	fmt.Println("\n Результат:")
	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)
}
