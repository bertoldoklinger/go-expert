package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10, 12}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:3]), cap(s[:3]), s[:3])
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 12)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
}
// quando colocamos o :, ele cria um slice de um slice, se for antes, só pega os numeros deposi do idndice, se for depois pega antes do indice

// por baixo dos panos, toda vez que tu da um append no slice, e ele não tem essa capacidade
// por debaixo dos panos ele vai dobrar a capacidade desse slice para alocar tudo e vai jogar nesse array novo

//para trabalhar com muitos dados em um slice, iniciar o slice num valor proximo que voce acha que vai trabalhar
// se não vai ser mais critico á performance do go
