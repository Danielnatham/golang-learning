package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Println("Hello", name)
}

func sayBye(name string) {
	fmt.Println("Bye", name)
}

func cycleNames(names []string, function func(string)) {
	for _, name := range names {
		function(name)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// name := "Mario"

	// sayGreeting(name)
	// sayBye(name)

	names := []string{"Jose", "Maria", "Carlos"}

	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)

	fmt.Printf("%0.3f", circleArea(1)*circleArea(2))
}
