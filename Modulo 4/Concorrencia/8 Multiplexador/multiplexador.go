package main

import (
	"fmt"
	"math/rand"
	"time"
)

// PARA AGRUPAR UM OI MAIS CANAIS
func mainD() {

	canalAgrupado := multiplexar(escreverTexto("Canal 1", "Teste"), escreverTexto("Canal2", "Teste"))

	for i := 0; i < 10; i++ {
		fmt.Println("Resultado do canal ", <-canalAgrupado)
	}

}

func multiplexar(canal1, canal2 <-chan string) chan string {
	canalDeSaida := make(chan string)

	go func() {

		for {

			select {
			case mensagem := <-canal1:
				canalDeSaida <- mensagem
			case mensagem := <-canal2:
				canalDeSaida <- mensagem
			}
		}

	}()

	return canalDeSaida
}

func escreverTexto(titulo string, texto string) <-chan string {
	canal := make(chan string)

	var i int = 0

	go func() {

		for {
			canal <- fmt.Sprintf("%v %v %v", titulo, texto, i+1)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
			i++

		}
	}()

	return canal

}
