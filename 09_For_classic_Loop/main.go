package main

import "fmt"

func sumOfNNumber() {
	N := 30
	sum := 0
	for i := 0; i <= N; i++ {
		sum += N
	}
	fmt.Println(sum)
}

func main() {
	sumOfNNumber()
}
