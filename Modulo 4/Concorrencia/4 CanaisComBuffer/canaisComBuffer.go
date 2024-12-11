package main

import "fmt"

func mainD() {

	//Nessa declaração diz que aceita apenas dois valores
	canal := make(chan string, 2)

	canal <- "teste do primeiro item"
	canal <- "teste do segundo item "

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
