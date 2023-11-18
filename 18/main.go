package main

import (
	"fmt"
	"curso-go/matematica"
)

func main() {

	s := matematica.Soma(10,20)
	
	fmt.Printf("Resultado: %v", s)
}