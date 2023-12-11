package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	valor, err := sum(50, 10)
	if err != nil {
		println(err)
	}
	fmt.Println(valor)

	fmt.Println(sum(51, 2))

	fmt.Println(divide(10, 2))

	strValor, err := toUppercase("r")

	fmt.Println(strValor)

	if err != nil {
		fmt.Println(err)
	}
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("a soma é maior que 50")
	}
	return a + b, nil
}

func divide(a, b int) int {
	return a / b
}

func toUppercase(name string) (string, error) {
	if len(name) < 2 {
		return "", errors.New("erro")
	}

	return strings.ToUpper(name), nil
}
