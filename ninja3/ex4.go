package main

import "fmt"

func main() {
	x := 2002

	for {
		if x > 2025 {
			break
		} else {
			fmt.Println(x)
			x++
		}
	}
}
