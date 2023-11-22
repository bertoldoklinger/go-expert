package main

func main() {
	// condicionais
	a := 1
	b := 2
	c := 3

	if a > b && c > a {
		println("a > b && c > a")
	}

	if a > b {
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("b")
	default:
		println("d")
	}
}
