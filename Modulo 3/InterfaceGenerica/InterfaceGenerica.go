package main

import "fmt"

func iGenerica(as interface{}) {

	fmt.Println(as)
}
func mainD() {
	iGenerica(10)
	iGenerica("teste")
	iGenerica(true)
	iGenerica(10.9)

}
