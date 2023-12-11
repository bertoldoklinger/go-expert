package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"number"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"Numero":2, "Saldo": 200}`)

	var contaX Conta

	err = json.Unmarshal(jsonPuro, &contaX)

	if err != nil {
		println(err)
	}

	println(contaX.Saldo)
}
