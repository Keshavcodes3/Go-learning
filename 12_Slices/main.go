package main

import "fmt"

func Array() {
	results := []string{"Keshav", "Sree"}
	otherNames := []string{"Atishay", "Anurag", "Vaivbhav"}
	results = append(results, otherNames...)
	fmt.Println(results)
}

func CapacityAndLen() {
	//len : How many you have
	//Capacity : How many you can append
	scores := make([]int, 0, 5)
	fmt.Println(scores,len(scores),cap(scores))
}

func main() {

	CapacityAndLen()
}
