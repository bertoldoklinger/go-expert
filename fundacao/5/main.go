package main

import "fmt"

func main() {

	var meuArray [4]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	meuArray[3] = 40

	for indice, valor := range meuArray {
		fmt.Printf("O meu indice é %d e o valor é %d\n", indice, valor)
	}
}

//no fmt string coloca %v e digito coloca %d
