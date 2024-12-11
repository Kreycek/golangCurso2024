package main

import "fmt"

func mainD() {

	var variavel int = 10
	var variavel2 int = variavel

	variavel2++

	fmt.Println(variavel, variavel2)

	var variavel3 int = 100
	var ponteiro *int

	ponteiro = &variavel3 //Pega o endereço de memória
	fmt.Println(variavel3, ponteiro)

	variavel3 = 150

	fmt.Println(variavel3, *ponteiro) //imprime o valor do ponteiro

	/*
		Use ponteiros para permitir que uma função modifique o valor original de uma variável.
		Use ponteiros para evitar cópias grandes de dados, especialmente com structs grandes.
		Use ponteiros quando você precisar representar a ausência de valor (nil).
		Use ponteiros para manipular diretamente estruturas de dados complexas, como listas encadeadas e árvores.
		Em interfaces, use ponteiros para garantir que os métodos possam modificar o estado interno de uma struct.
	*/
}
