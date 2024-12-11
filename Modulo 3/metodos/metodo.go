package main

import "fmt"

func fnc() int64 {

	return 50 - 10
}
func mainD() {

	t := fnc()

	fmt.Println(t)

	x, o := fnc2()

	fmt.Println(x, o)

	a, _ := fnc2()

	fmt.Println(a)
}

func fnc2() (int, int) {

	return 10, 10
}
