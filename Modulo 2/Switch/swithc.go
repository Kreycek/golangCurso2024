package main

import "fmt"

func mainD() {
	for i := 10; i > 1; i-- {

		switch i {
		case 7:
			fmt.Println("É ", i+1)
			break
		case 5:
			fmt.Println("É ", i+1)
			break
		}
		fmt.Println(i)
	}
	// var u int =10
}
