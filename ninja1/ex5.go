package main

import "fmt"

type valorant int
var x valorant
var y int

func main() {
	
	fmt.Printf("%v, %T\n", x, x)

	x = 42

	fmt.Println(x)

	y = int(x)

	fmt.Printf("%v, %T\n", y, y)

}