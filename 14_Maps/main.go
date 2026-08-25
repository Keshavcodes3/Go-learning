package main

import "fmt"

func NilMap() {
	var scores map[string]int
	scores = make(map[string]int)
	scores["Keshav"] = 90
	fmt.Println(scores["Keshav"])
}

func Map() {
	var months map[string]int
	months = make(map[string]int)
	months["January"] = 30
	for i, v := range months {
		fmt.Println(v, i)
	}
}

func Map2() {
	var scores map[string]int
	scores = make(map[string]int)
	scores["Keshav"] = 20
	scores["Anurag"] = 58
	scores["Sree"] = 90
	fmt.Println(scores)

	delete(scores, "Keshav")

	fmt.Println(scores)

}

func main() {
	//map[keyType]valueType
	// ages := map[string]int{
	// 	"Keshav": 18,
	// 	"John":   35,
	// }
	// fmt.Println(ages["Keshav"])
	// NilMap()
	Map2()
}
