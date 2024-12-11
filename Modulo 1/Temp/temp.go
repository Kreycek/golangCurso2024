package main

import "fmt"

const ebulicaoF float64 = 212.0

var tt int = 10

func mainD() {
	var tempF float64 = ebulicaoF
	var tempC float64 = (tempF - 32) * 5 / 9

	fmt.Println("A temperatura de ebulição da água em F = ", tempF)
	fmt.Println("A temperatura de ebulição da água em C = ", tempC)

	fmt.Printf("A temperatura de ebulição da água em C é = %g e em F é %g ", tempC, tempF)

}
