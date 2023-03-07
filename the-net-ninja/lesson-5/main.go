package main

import "fmt"

func main() {
	// loop()
	// loopIncrement()
	// loopArray()
	loopArrayRange()
}

func loop() {
	x := 0

	for x <= 5 {
		fmt.Println(x)

		x++
	}
}

func loopIncrement() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}

func loopArray() {
	names := []string{"Dan", "Victor", "Lucas"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

func loopArrayRange() {
	names := []string{"Dan", "Victor", "Lucas"}

	// for index, value := range names {
	// 	fmt.Printf("Index: %v Value: %v \n", index, value)
	// }

	//Without index
	for _, value := range names {
		fmt.Printf("Value: %v \n", value)
	}
}
