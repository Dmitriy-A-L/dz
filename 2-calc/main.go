package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	if len(nums) == 0 {
		fmt.Println("Список чисел пуст.")
		return
	}
	switch strings.ToUpper(op) {
	case "SUM":
		fmt.Printf("Сумма: %.2f\n", sum(nums))
	case "AVG":
		fmt.Printf("Среднее: %.2f\n", avg(nums))
	case "MED":
		fmt.Printf("Медиана: %.2f\n", median(nums))
	default:
		fmt.Println("Неизвестная операция. Доступны: AVG, SUM, MED.")
	}
}
func sum(nums []float64) float64 {
	total := 0.0
	for _, n := range nums {
		total += n
	}
	return total
}
func avg(nums []float64) float64 {
	return sum(nums) / float64(len(nums))
}
func median(nums []float64) float64 {
	sort.Float64s(nums)
	n := len(nums)
	mid := n / 2
	if n%2 == 0 {
		return (nums[mid-1] + nums[mid]) / 2
	}
	return nums[mid]
}
