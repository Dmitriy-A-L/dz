package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var op string
	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	fmt.Scan(&op)
	var input string
	fmt.Print("Введите числа через запятую (например: 2,10,9): ")
	fmt.Scan(&input)
	bufio.NewReader(os.Stdin)
	parts := strings.Split(input, ",")
	nums := make([]float64, 0, len(parts))
	for _, p := range parts {
		n, err := strconv.ParseFloat(strings.TrimSpace(p), 64)
		if err != nil {
			fmt.Println("Ошибка: введите только числа, разделенные запятыми.")
			return
		}
		nums = append(nums, n)
	}
}
