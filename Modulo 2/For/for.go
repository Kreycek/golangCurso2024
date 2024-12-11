package main

import "fmt"

func mainD() {
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i > 10; i-- {
		fmt.Println(i)
	}

}
