package main

func main() {
	//type assertion
	var minhaVar interface{} = "Bertoldo Klinger"

	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	println(res, ok)
	res2 := minhaVar.(int)
	println(res2)
}
