package main

import "fmt"

func main() {
	//interfaces vazias
	var x interface{} = 10
	var y interface{} = "Hello world"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é  %T, e o valor é %v", t, t)
}
