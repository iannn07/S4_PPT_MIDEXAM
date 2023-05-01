package main

import "fmt"

func main() {
	var a [3]string
	var b [3]string

	fmt.Println("Array 1:")
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println("Array 2:")
	for i := 0; i < len(b); i++ {
		fmt.Scan(&b[i])
	}

	compare := comparation(a, b)
	if len(compare) > 0 {
		for _, index := range compare {
			fmt.Printf("Index-%d is different\n", index)
		}
	} else {
		fmt.Println("Arrays are equal")
	}
}

func comparation(a [3]string, b [3]string) []int {
	var temp []int
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			temp = append(temp, i)
		}
	}
	return temp
}
