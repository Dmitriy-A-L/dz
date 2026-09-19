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

// есть ли такая валюта в списке
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
		// чтение, игнориирование ошибки
		input, _ := reader.ReadString('\n')
		// затираем пробелы и делаем заглавными
		input = strings.TrimSpace(strings.ToUpper(input))
		// если валюта корректна - выводим
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

// подсчёт сумма из стоки в строку, дробное.
// func calculate(amount float64, from string, to string) float64 {
// 	// курсы валют (с map не понял)
// 	rates := map[string]float64{
// 		"USD": 90.0,
// 		"EUR": 100.0,
// 		"RUB": 1.0,
// 	}

// 	inRub := amount * rates[from]
// 	result := inRub / rates[to]
// 	return result
// }

// вход, ввод исходной валюты
func main() {
	const usdtoEur = 0.871
	const usdtoRub = 84.1975
	const eurtoRub = 96.6671
	const rubtoEur = 0.0103
	const eurtoUsd = 1.1481
	const rubtoUsd = 0.0119
	fmt.Println("💰 Калькулятор валют")
	fmt.Println("\nШаг 1: Выберите исходную валюту")
	from := inputCurrency("Исходная валюта")
	// ввод числа
	fmt.Println("\nШаг 2: Введите сумму")
	amount := inputNumber("Сумма: ")
	// ввод целевой валюты
retryTo:
	fmt.Println("\nШаг 3: Выберите целевую валюту")
	to := inputCurrency("Целевая валюта")
	if from == to {
		fmt.Println("\n Валюта совпадает с исходной, введите другую валюту")
		goto retryTo
	}
	// задаём переменную результата
	var result float64
	// запускаем условия
	switch {
	// ввод баксов и вывод евры
	case from == "USD" && to == "EUR":
		result = amount * usdtoEur
		// etc
	case from == "USD" && to == "RUB":
		result = amount * usdtoRub
	case from == "EUR" && to == "USD":
		result = amount * eurtoUsd
	case from == "EUR" && to == "RUB":
		result = amount * eurtoRub
	case from == "RUB" && to == "USD":
		result = amount * rubtoUsd
	case from == "RUB" && to == "EUR":
		result = amount * rubtoEur
	}
	// пишем результат
	fmt.Println("\n Результат:")
	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)

}
