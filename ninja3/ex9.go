package main

import "fmt"

func main() {
	esporteFavorito := "Futebol"

	switch esporteFavorito {
	case "Futebol":
		fmt.Println("Futebol")
	case "Basquete":
		fmt.Println("Basquete")
	case "Volei":
		fmt.Println("Volei")
	default:
		fmt.Println("Deu em nada")
	}
}
