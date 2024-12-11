package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("recuperada execução")
	}
}

func mainD() {

	defer recuperarExecucao()

	fmt.Println("Teste 1")

	panic("ok")
}
