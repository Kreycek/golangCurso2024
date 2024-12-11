package main

import (
	"fmt"
	"time"
)

func mainD() {

	n := 0
	for n < 2 {
		n++
		fmt.Println(n)
		time.Sleep(2000)
	}

	for j := 1; j < 3; j++ {
		fmt.Println(j)
		time.Sleep(1000)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	fmt.Println("")

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("")
	for indice, nome := range "PALAVRA" {
		fmt.Println(indice, string(nome))
	}

	fmt.Println("")
	fmt.Println("")

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Isso abaixo é um loop infinito
	for {
		break
	}
}
