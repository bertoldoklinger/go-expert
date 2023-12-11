package main

import "fmt"

//composicao de struct
//structs são tipos de classes, não classes

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}
// interfaces só possuem assinatura de métodos, não de props
// quando usamos interface, automaticamente QUALQUER struct que tiver o método Desativar() estará implementando a interface Pessoa, como um implements do OOP
type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

// em métodos de struct, para ele se tornar um método, abrimos parantese antes do nome da function e atachamos a struct com um alias antes
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf(c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	bertoldo := Cliente{
		Nome:  "Bertoldo",
		Idade: 24,
		Ativo: true,
	}
	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)

	fmt.Printf("Nome: %v, Idade: %d, Ativo: %t", bertoldo.Nome, bertoldo.Idade, bertoldo.Ativo)
}
