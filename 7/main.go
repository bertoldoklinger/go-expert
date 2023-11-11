package main

import "fmt"

//maps também são chamados de hashtables
// estruturas de dados key value

func main() {

	salarios := map[string]int{"Bertoldo": 1000, "João": 2000, "Maria": 3000}
	delete(salarios, "Maria")
	salarios["Marcos"] = 5000

	// sal := make(map[string]int)
	// sal2 := map[string]int {}

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}
	for _, salario := range salarios {
		fmt.Println(salario)
	}

	// para ignorarmos um valor em um loop, podemos fazer a annotation _ (blank identifier)

}
