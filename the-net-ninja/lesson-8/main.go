package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	str := strings.ToUpper(name)

	names := strings.Split(str, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], ""
}

func main() {
	firstInitial, secondInitial := getInitials("Daniel Nathan")

	fmt.Println(firstInitial, secondInitial)

	firstInitial1, secondInitial1 := getInitials("Logan")

	fmt.Println(firstInitial1, secondInitial1)
}
