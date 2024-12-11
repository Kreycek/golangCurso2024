package main

import (
	"fmt"
)

func soma(texto string, numeros ...int) (string, int) {

	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return texto, total
}

func mainD() {

	texto, totalDaSoma := soma("Total: ", 1, 4, 5, 6, 78)
	fmt.Println(texto, totalDaSoma)
}
