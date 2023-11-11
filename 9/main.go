package main

import "fmt"

//funcoes variaticas, ...int

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16))
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
