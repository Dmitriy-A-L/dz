package main

import "fmt"

func main() {
	const usdToEur = 0.92
	const usdToRub = 97.50
	const eurToRub = usdToRub / usdToEur

	usdAmount := 100.0
	eurAmount := usdAmount * usdToEur
	rubAmount := usdAmount * usdToRub
	eurToRubAmount := eurAmount * eurToRub

	fmt.Printf("%.2f USD = %.2f EUR\n", usdAmount, eurAmount)
	fmt.Printf("%.2f USD = %.2f RUB\n", usdAmount, rubAmount)
	fmt.Printf("%.2f EUR = %.2f RUB\n", eurAmount, eurToRubAmount)
	fmt.Printf("Курс EUR/RUB: %.2f\n", eurToRub)
}
