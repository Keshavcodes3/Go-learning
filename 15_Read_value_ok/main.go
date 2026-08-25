package main

import "fmt"

func main() {
	points := map[string]int{
		"a": 20,
		"b": 52,
		"c": 55,
	}
	for idx, val := range points {
		fmt.Println(idx, val)
	}
	
}
