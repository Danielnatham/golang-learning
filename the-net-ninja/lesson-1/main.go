package main

import "fmt"

var nameOne string = "Escopo"

func main()  {
	//strings
	var nameOne string = "Mario"
	var nameTwo = "Luigi"
	var nameThree string
	nameFour := "Bowser"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	//numbers
	var ageOne int = 30
	var ageTwo = 20
	ageThree := 10

	println(ageOne, ageTwo, ageThree)

	//bits
	var numberOne int8 = -124
	var numberTwo uint = 255

	println(numberOne, numberTwo)

	//floats

	var scoreOne float32 = 25.40
	var scoreTwo float64 = -30.404030

	println(scoreOne, scoreTwo)
}