package main

import (
	"curso-go/matematica"
	"fmt"

	"github.com/google/uuid"
)

func main() {

	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Printf("Resultado: %v", s)
	fmt.Println(matematica.A)
	fmt.Println(carro.Marca)
	matematica.Andar(&carro)
	fmt.Println(uuid.New())
}
