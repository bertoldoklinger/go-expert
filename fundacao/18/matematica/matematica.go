package matematica

import "fmt"

func Soma[T int | float64](a T, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func Andar(c *Carro) {
	fmt.Printf("Carro andando %v", c)
}
