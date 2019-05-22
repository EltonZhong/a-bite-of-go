package main

import "fmt"

var block = "1"

func main() {
	fmt.Print(block)

	block = "12"
	fmt.Println(block)

	block := 12
	fmt.Println(block)

	block = 1
	fmt.Println(block)
	{
		var block []int = nil
		block = []int {1, 2, 3}
		fmt.Println(block)
	}

	var jj = 1
	fmt.Println(string(jj))
	fmt.Println(string(jj))
	fmt.Println(string(jj))
	fmt.Println(string(jj))
	fmt.Println(string(122222))

	type DamnString = string

	type GodDamnString string
}
