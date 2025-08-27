package main

import "fmt"

type valorant int
var x valorant

func main() {
	
	fmt.Printf("%v, %T\n", x, x)

	x = 42

	fmt.Println(x)

}