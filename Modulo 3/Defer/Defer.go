package main

import (
	"fmt"
)

func funcao1() {

	fmt.Println("Executando a função 1 ")
}

func funcao2() {

	defer fmt.Println("Executando a função 2 defer ")

	fmt.Println("Executando a função 2 ")
}

func mainD() {

	defer funcao1() // Deixa essa execução por último
	funcao2()
	fmt.Println("teste")
}
