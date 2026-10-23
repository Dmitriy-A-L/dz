package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Список доступных валют
var availableCurrencies = []string{"USD", "EUR", "RUB"}
var ratesPerUSD = map[string]float64{
	"USD": 1.0,
	"EUR": 0.871,
	"RUB": 84.1975,
}

func isValidCurrency(currency string) bool {
	// запуск цикла для проверки введённой валюты
	for _, c := range availableCurrencies {
		// при совпадении возврат и завершение, иначе ошибка

		if c == currency {
			return true
		}
	}
	return false
}

// функция ввод валюты по типу строка
func inputCurrency(prompt string) string {
	// чтение стандартного значения с клавиатуры
	reader := bufio.NewReader(os.Stdin)
	// запуск бесконечного цикла
	for {
		// вставляем строку доступных валют
		fmt.Printf("%s (доступно: %s): ", prompt, strings.Join(availableCurrencies, ", "))
		// чтение, игнорирование ошибки
		input, _ := reader.ReadString('\n')
		// затираем пробелы и делаем заглавными
		input = strings.TrimSpace(strings.ToUpper(input))
		// если валюта корректна - вводим
		if isValidCurrency(input) {
			return input
		}
		// иначе выводим ошибку
		fmt.Println("❌ Ошибка: такой валюты нет. Попробуйте снова.")
	}
}

// вводим дробное число
func inputNumber(prompt string) float64 {
	// читаем
	reader := bufio.NewReader(os.Stdin)
	// бесконечный цикл
	for {
		// печать, чтение с игнором ошибки, с новой строки, убираем пробелы
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		// перевод строки в число, вывод числа и ошибки
		number, err := strconv.ParseFloat(input, 64)
		// если нет ошибки И число больше 0, то выводим число
		if err == nil && number > 0 {
			return number
		}
		// иначе вывод сообщения
		fmt.Println("❌ Ошибка: введите корректное положительное число.")
	}
}

// вход, ввод исходной валюты
func main() {

	fmt.Println("💰 Калькулятор валют")
	fmt.Println("\nШаг 1: Выберите исходную валюту")
	from := inputCurrency("Исходная валюта")
	// ввод числа
	fmt.Println("\nШаг 2: Введите сумму")
	amount := inputNumber("Сумма: ")
	// ввод целевой валюты
	fmt.Println("\nШаг 3: Выберите целевую валюту")
	var to string
	for {
		to = inputCurrency("Целевая валюта")
		if to != from {
			break
		}
		fmt.Println("Валюта совпадает с исходной, введите другую")
	}

	amountInUSD := amount / ratesPerUSD[from]
	result := amountInUSD * ratesPerUSD[to]

	// пишем результат
	fmt.Println("\n Результат:")
	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)

}
