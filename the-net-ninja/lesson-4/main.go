package main

import (
	"fmt"
	"sort"
)

func main() {

	// greeting := "Hello World"
	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Split(greeting, " "))
	// fmt.Println("Original string: ", greeting)

	ages := []int{1, 2, 5, 23, 10, 3, 50, 21}
	names := []string{"Dan", "Baniel", "Anilo"}
	sort.Ints(ages)
	sort.Strings(names)

	fmt.Println(ages)

	fmt.Println(sort.SearchStrings(names, "Danilo"))

}
