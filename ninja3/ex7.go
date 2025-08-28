package main

import "fmt"

func main() {
	x := 14

	if x == 10 {
		fmt.Println("X vale 10")
	} else if x < 10 {
		fmt.Println("X vale menos que 10")
	} else {
		fmt.Println("X vale mais que 10")
	}
}
