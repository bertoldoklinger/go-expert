package main

func main() {
	// Memória -> Endereço -> Valor

	a := 10
	// quando usamos &a, pegamos o endereço onde a variável está apontando
	// quando usamos *, estamos falando do endereço que está na memória, não do valor
	// Variável -> Ponteiro que tem um endereço na memória -> Valor
	var ponteiro *int = &a
	*ponteiro = 20
	println(a)
	b := &a
	// * é derreferencing, ele sabe que tá apontando para um valor na memória,quando colocamos * ele pergunta qual valor está guardado na memória
	*b = 30
	println(*b)
	println(a)

	partialPrice := 20

	var ponteiro2 *int = &partialPrice

	*ponteiro2 = 20

	println(ponteiro2, &partialPrice)

	totalPrice := &partialPrice

	*totalPrice = 50

	println(*totalPrice, partialPrice)
}
