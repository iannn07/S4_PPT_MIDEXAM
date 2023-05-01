package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("This code is written by - Pristian Budi Dharmawan")
	fmt.Println("Ellipsis - Problem 1")

	// Ellipsis is a method to make a function or a method can accept variable elements or in easy words its not limited to passing variable elements as many as you want.

	fmt.Println("\nEllipsis Declaration")
	q := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%T\n", q)

	fmt.Println("\nEllipsis Example (1)")
	a := [3]string{"a", "b", "c"}
	b := [...]string{"a", "b", "c"}
	fmt.Println("Does array A & B equals?", a == b)

	fmt.Println("\nEllipsis Example (2)")
	fmt.Printf("The total of 10 + 30 + 20 + 40 = ")
	seq1 := example2(10, 30, 20, 40)
	fmt.Println(seq1)

	fmt.Println("\nEllipsis Example (3)")
	fmt.Println(example3("Hello", "World!"))
}

func example2(seq ...int) int {
	total := 0
	for _, num := range seq {
		total += num
	}
	return total
}

func example3(str ...string) string {
	return strings.Join(str, " ")
}
