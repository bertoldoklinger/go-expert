package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Bertoldo": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Bertoldo": 100.0, "João": 200.0, "Maria": 300.0}
	m3 := map[string]MyNumber{"Bertoldo": 100.0, "João": 200.0, "Maria": 300.0}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}
