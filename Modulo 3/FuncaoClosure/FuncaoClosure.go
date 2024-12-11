package main

import "fmt"

func closure() func() {
	texto := "dentro da função closure"
	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func mainD() {

	funcaoNova := closure()

	funcaoNova()

}
