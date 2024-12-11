package main

import (
	"fmt"
)

// Função recursiva para verificar se o número é 10
func verificaSeDez(n int) bool {
	if n == 10 {
		return true
	}
	if n < 10 {
		return false
	}

	fmt.Println(n)
	return verificaSeDez(n - 1)
}

func mainD() {
	num := 15
	if verificaSeDez(num) {
		fmt.Printf("%d eventualmente chega a 10\n", num)
	} else {
		fmt.Printf("%d não é ou não chega a 10\n", num)
	}
}
