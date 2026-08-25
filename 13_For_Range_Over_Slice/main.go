package main

import "fmt"

func Range() {
	months := make([]string, 0, 10)
	months = append(months, "January", "February", "March", "April", "May")
	for idx, val := range months {
		fmt.Println(idx+1, "->", val)
	}
}

func main() {
	// views := []int{10, 20, 30, 40, 50, 60}

	// total := 0
	// for idx, val := range views {
	// 	fmt.Println("day ", idx, "views", val)
	// 	total += val
	// }
	// fmt.Println("total ", total)
	Range()
}
