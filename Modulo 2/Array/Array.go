package main

import (
	"fmt"
	"strconv"
)

func mainD() {

	var nn []string
	nn = append(nn, "Lucas")
	var notas [5]float64

	notas[0] = 6.7
	notas[1] = 7.1
	notas[2] = 5.2
	notas[3] = 4.5
	notas[4] = 3.3

	var total float64 = 0

	for i := 0; i < len(notas); i++ {

		total += notas[i]

	}

	fmt.Println("A média é " + strconv.FormatFloat(total, 'f', 2, 64))

}
