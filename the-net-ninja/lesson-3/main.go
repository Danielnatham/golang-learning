package main

import "fmt"

func main() {
	// Array
	var ages = [3]int{20, 25, 35}
	names := [4]string{"Browser", "Mario", "Peach", "Yos"}

	names[0] = "Luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slice
	var scores = []int{100, 50, 60}
	scores[2] = 25

	scores = append(scores, 10)

	fmt.Println(scores, len(scores))

	// Slice ranges
	rangerOne := names[1:2]
	rangerTwo := names[2:]
	
	fmt.Println(rangerOne, rangerTwo)
}