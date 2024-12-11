package main

import "fmt"

type pessoa struct {
	nome     string
	endereco string
}

func mainD() {

	var p pessoa

	p.nome = "Ricardo Silva "
	p.endereco = "Rua Lima"

	fmt.Println(p)
}
