package main

import "fmt"

func mainD() {

	fmt.Print("teste ")

	fmt.Println(teste(1))

	fmt.Println(teste2_fallthrough(1))
}

func teste(numero int32) string {

	switch numero {
	case 1:
		return "OK"
		fallthrough
		//Continuar para o pr+oximo case mesmo que seja verdadeira a condição
		//Se cair nesse case vai direto para o case 2
	case 2:
		return "NOK"
	default:
		return "vazio"
	}

}

func teste2_fallthrough(numero int32) string {

	var diaSemana string

	switch {
	case numero == 1:
		diaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaSemana = "Segunda"
	default:
		diaSemana = "vazio"
	}

	return diaSemana
}
