package main

import "fmt"

type pessoa struct {
	nome     string
	endereco string
}

type estudante struct {
	pessoa
	nota1 int
}

func mainD() {

	var est estudante
	var p pessoa = pessoa{"Ricardo", "Rua 7"}

	est = estudante{p, 1}

	est = estudante{pessoa{"Ricardo", "Rua 7"}, 1}

	fmt.Println(est)

	fmt.Printf(est.nome)
}
