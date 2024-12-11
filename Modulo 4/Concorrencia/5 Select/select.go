package main

import (
	"fmt"
	"time"
)

func mainD() {
	canal1, canal2 := make(chan string), make(chan string)
	contador := 0
	limite := 10 // Número máximo de mensagens que queremos processar

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
			contador++

		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
			contador++
		}

		if contador >= limite {
			fmt.Println("Limite atingido, encerrando o programa.")
			close(canal1)
			close(canal2)
			break // Encerra o programa
		}
	}
}
