package main

import "fmt"

func main() {
	x := 4

	switch {
	case x == 3:
		fmt.Println("O numero escolhido foi o 3")
	default:
		fmt.Println("Deu em nada")
	}
}
