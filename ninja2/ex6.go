package main

import "fmt"

const (
	_ = iota + 2025
	a
	b
	c
	d
)

func main() {
	fmt.Printf("The next 4 years will be:\n%v\n%v\n%v\n%v\n", a, b, c, d)
}
