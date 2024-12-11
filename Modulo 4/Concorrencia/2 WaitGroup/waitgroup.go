package main

import (
	"fmt"
	"sync"
	"time"
)

func mainD() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(3)
	go func() {
		escrever("GO Rotine 1")
		waitGroup.Done()
	}()

	go func() {
		escrever("GO Rotine 2")
		waitGroup.Done()
	}()

	go func() {
		escrever("GO Rotine 3")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
