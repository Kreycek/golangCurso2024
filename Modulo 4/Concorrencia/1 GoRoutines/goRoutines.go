package main

import (
	"fmt"
	"time"
)

func mainD() {
	go escrever("Olá mundo")
	escrever("programado em GO")

	fmt.Println("tere")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
