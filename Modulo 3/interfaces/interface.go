package main

import (
	"fmt"
	"strconv"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func soma(valor1 int, valor2 int) int {

	return valor1 + valor2
}

func subtracao(valor1 int64, valor2 int64) int64 {

	return valor1 - valor2
}

type operacoes interface {
	soma(valor1 int, valor2 int) int
	subtracao(valor1 int64, valor2 int64) int64
}

func mainD() {

	var g operacoes

	result := g.soma(30, 90)

	fmt.Printf(strconv.Itoa(result))

	var i int64 = 10

	o := 90

	fmt.Println(i, o)

	r := retangulo{10, 15}

	escreverArea(r)
}
