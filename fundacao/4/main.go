package main

import (
	"fmt"
)

type ID int

var (
	f ID = 123
)

func main() {

	fmt.Printf("O tipo de E é %T\n", f)
}
