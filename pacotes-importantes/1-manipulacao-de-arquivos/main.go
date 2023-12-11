package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso o tamanho é: %v", tamanho)

	f.Close()
	/// leitura

	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	//leitura de pouco em pouco abrindo o arquivo

	arquivo2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(arquivo2)

}
