package main

import (
	"fmt"
)

func soma(n1, n2 int) (soma int, subtracao int) {

	soma = n1 + n2
	subtracao = n1 - n2

	return //aqui ja entende que é para retornar a variavel soma e subtração
}

func mainD() {

	m1, m2 := soma(10, 20)
	fmt.Println(m1, m2)
	fmt.Println("teste")
	fmt.Println(soma(10, 40))

}
