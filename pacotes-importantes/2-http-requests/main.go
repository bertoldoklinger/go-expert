package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://api-pdt.vercel.app/")

	if err != nil {
		panic(err)
	}

	result, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	println(string(result))
	req.Body.Close()
}
