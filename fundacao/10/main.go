package main

//clojures ou funções anonimas que se auto executam

import "fmt"

func main() {

	func() {
	}()

	total := func() int {
		return sum(1, 2, 3, 5, 4, 6, 7, 8, 9) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
