package main

import "fmt"

type Carro struct {
	Marca       string
	ReleaseYear int
}

func main() {
	// loop padrao
	for i := 0; i < 10; i++ {
		println(i)
	}
	// loopando sobre objetos , loop mais facil
	carros := []Carro{{"Corolla", 1999}, {"Fiat", 1998}, {"Renault", 2007}}
	for i, v := range carros {
		if i > 1 {
			fmt.Println(v.ReleaseYear)
		}
	}
	//while
	i := 0
	for i < 10 {
		println(i)
		i++
	}
	// loop infinito
	for {
		println("'Hello world'")
	}
}
