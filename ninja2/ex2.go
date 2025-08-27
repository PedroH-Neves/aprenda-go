package main

import "fmt"

func main() {
	a := 7 == 70
	b := 7 != 70
	c := 7 >= 70
	d := 7 > 70
	e := 7 <= 70
	f := 7 < 70

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}
