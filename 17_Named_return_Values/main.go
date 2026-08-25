package main

import "fmt"

func divideReminderSumProd(a int, b int, c int) (sum int, prod int, divide int, rem int) {
	sum = a + b + c
	prod = a * b * c
	divide = a / b
	rem = a / c
	return sum, prod, rem, divide
}

func main() {
	prod, _, _, _ := divideReminderSumProd(3, 3, 3)
	fmt.Println(prod)
}
