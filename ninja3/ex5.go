package main

import "fmt"

func main() {
	//mostrar o resto da divisao por 4 de numeros de 10 a 100
	for x := 10; x <= 100; x++ {
		y := x % 4

		fmt.Println(y)
	}
}
