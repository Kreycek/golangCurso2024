package main

import "fmt"

type Person struct {
	nome      string
	sobrenome string
	idade     int64
}

func mainD() {

	var dd string = "Maria"
	fmt.Printf(dd)
	fmt.Printf("teste")
	tt := Person{" Maria", " Santos Silva", 56}

	fmt.Print(tt.nome)
}
