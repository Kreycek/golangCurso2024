package main

import (
	"fmt"
	"strconv"
	"time"
)

func mainD() {

	canal := make(chan string)
	go escrever("Olá Mundo ", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// for {
	// 	mensagem, aberto := <-canal

	// 	if !aberto {
	// 		break
	// 	}

	// 	fmt.Println(mensagem)
	// }

	//Esse código abaixo substitui o acima
	for mensagem := range canal {

		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {

		canal <- texto + strconv.Itoa(i+1)
		// fmt.Println(texto)
		time.Sleep(time.Second)
	}

	close(canal)
}
