package main

import "fmt"

func init() {
	//Essa aqui é executada primeiro idela por exemplo para inicializar variáveis globais
	fmt.Println("Teste main")
}

func mainD() {

	fmt.Println("func main")
}
