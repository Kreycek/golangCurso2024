package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (s usuario) anos() {
	fmt.Println(" A idade é ", 18)

}

func mainD() {

	var usuar usuario

	usuar.nome = "Ricardo"
	usuar.idade = 20
	usuar.anos()

	fmt.Println(" A idade é ", usuar.nome)

}
