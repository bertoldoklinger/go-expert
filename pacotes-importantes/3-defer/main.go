package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda linha")
	defer fmt.Println("Terceira linha")
}