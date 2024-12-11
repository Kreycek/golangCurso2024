package main

import (
	"fmt"
)

func mainD() {

	var ed = func(text string) string {
		return text
	}("texto")
	fmt.Println("ok", ed)

}
