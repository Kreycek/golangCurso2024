package main

import (
	"fmt"
)

type nnn struct {
	nome  string
	idade int
	preco []float32
}

func mainD() {
	var teste nnn
	teste.nome = "Ricardo Silva"
	teste.idade = 40
	teste.preco = append(teste.preco, 10.6)
	teste.preco = append(teste.preco, 20.19)

	var nome string = "Tenille"
	var idade int = 36
	var versao float32 = 2.2
	fmt.Println("Olá Sr.", nome, " sua idade é ", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("valor da estrutura", teste)

}
