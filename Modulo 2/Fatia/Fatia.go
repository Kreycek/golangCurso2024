package main

import "fmt"

func mainD() {

	var nums = []int{1, 2, 3}

	nums = append(nums, 4, 5, 6)

	var num2 []int

	num2 = make([]int, 5)

	fmt.Println(nums)
	fmt.Println(num2)
	fmt.Println(num2[1:3])

	fatia1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fatia2 := make([]int, 10)

	copy(fatia2, fatia1)

	fmt.Println(fatia1, fatia2)
}
