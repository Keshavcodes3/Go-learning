package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func prodAndSum(a int, b int) (int, int) {
	return a + b, a * b
}

func main() {
	sum1 := add(2, 3)
	fmt.Println(sum1)
	_, product := prodAndSum(2, 3)
	fmt.Println(product)
}
