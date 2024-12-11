package main

import "fmt"

func teste(fdf *int) {

	fmt.Println(fdf)
}

func mainD() {

	var numero *int = new(int)

	*numero = 10

	fdf := 10

	teste(&fdf)

	fmt.Println(numero)

}
