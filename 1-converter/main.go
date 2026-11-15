package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var availableCurrencies = []string{"USD", "EUR", "RUB"}

var rates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.871,
	"RUB": 84.1975,
}

var operations = make(map[string]func(float64) float64)

var stdinReader = bufio.NewReader(os.Stdin)

func init() {
	for from := range rates {
		for to := range rates {
			if from != to {
				from, to := from, to
				operations[from+"_"+to] = func(amount float64) float64 {
					return (amount / rates[from]) * rates[to]
				}
			}
		}
	}
}

func isValidCurrency(currency string) bool {
	for _, c := range availableCurrencies {
		if c == currency {
			return true
		}
	}
	return false
}

func inputCurrency(prompt string) string {
	for {
		fmt.Printf("%s (доступно: %s): ", prompt, strings.Join(availableCurrencies, ", "))
		input, err := stdinReader.ReadString('\n')
		if err != nil {
			fmt.Println("\nВвод завершён.")
			os.Exit(0)
		}
		input = strings.TrimSpace(strings.ToUpper(input))
		if isValidCurrency(input) {
			return input
		}
		fmt.Println("Ошибка: такой валюты нет. Попробуйте снова.")
	}
}

func inputNumber(prompt string) float64 {
	for {
		fmt.Print(prompt)
		input, err := stdinReader.ReadString('\n')
		if err != nil {
			fmt.Println("\nВвод завершён.")
			os.Exit(0)
		}
		input = strings.TrimSpace(input)
		number, err := strconv.ParseFloat(input, 64)
		if err == nil && number > 0 {
			return number
		}
		fmt.Println("Ошибка: введите корректное положительное число.")
	}
}

func main() {
	fmt.Println("Калькулятор валют")

	fmt.Println("\nШаг 1: Выберите исходную валюту")
	from := inputCurrency("Исходная валюта")

	fmt.Println("\nШаг 2: Введите сумму")
	amount := inputNumber("Сумма: ")

	fmt.Println("\nШаг 3: Выберите целевую валюту")
	var to string
	for {
		to = inputCurrency("Целевая валюта")
		if to != from {
			break
		}
		fmt.Println("Валюта совпадает с исходной, введите другую")
	}

	operationKey := from + "_" + to
	convertFunc, exists := operations[operationKey]
	if !exists {
		fmt.Println("Операция для выбранной пары валют не найдена.")
		return
	}

	result := convertFunc(amount)
	fmt.Println("\nРезультат:")
	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)
}
