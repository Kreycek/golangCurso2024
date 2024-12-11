package main

import (
	"fmt"
	"time"
)

func mainD() {

	canal := escreverTexto("Olá mundo")

	for i := 0; i < 10; i++ {
		fmt.Println("Canal ", <-canal)
	}

	// fmt.Println(canal)
}

func escreverTexto(texto string) <-chan string {
	canal := make(chan string)

	var i int = 0

	go func() {

		for {
			canal <- fmt.Sprintf("Valor Recebido %v %v", texto, i+1)
			time.Sleep(time.Millisecond * 500)
			i++

		}
	}()

	return canal

}
