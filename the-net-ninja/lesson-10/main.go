package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":  4.50,
		"bread": 3.00,
		"salad": 1.50,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])
}
