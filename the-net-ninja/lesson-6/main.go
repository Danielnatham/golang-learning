package main

import (
	"fmt"
)

func main() {
	age := 50

	fmt.Println(age > 50)
	fmt.Println(age < 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 50 {
		fmt.Println("Age is less or equal 50")
	} else if age > 50 {
		fmt.Println("Age is less or equal 30")
	} else {
		fmt.Println("Age is 50")
	}

	names := []string{"Mario", "Luigi", "Peach", "Yo"}

	for index, name := range names {

		if index == 1 {
			fmt.Printf("Continue at %v \n", name)
		}

		if index > 2 {
			fmt.Printf("Break at %v", index)
			break
		}

		fmt.Printf("The value at pos %v is %v \n", index, name)

	}
}
