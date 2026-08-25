package main

import "fmt"

func ifWithShortStatement() {
	items := 3
	pricePerItem := 49
	if total := items * pricePerItem; total >= 100 {
		fmt.Println("Eligible for shopping")
	} else {
		fmt.Println("Not Eligible for shopping")
	}

}

func main() {
	score := 72
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else if score >= 45 {
		fmt.Println("C")
	} else {
		fmt.Println("Failed")
	}
	ifWithShortStatement()
}
