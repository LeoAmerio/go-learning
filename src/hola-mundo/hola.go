package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())

	// variables
	var name string = "Juan"
	// var age int = 20
	var isStudent bool = true
	var height, weight float64 = 1.75, 70.5
	var isMarried bool = false

	// constantes
	const pi float64 = 3.14
	const name string = "Juan"
	// const age int = 20
	const isStudent bool = true
	const height, weight float64 = 1.75, 70.5
	const isMarried bool = false

	// Inline declaration
	var firstName, lastName, age = "Juan", "Perez", 20
	fmt.Println("Nombre:", firstName, lastName, "Edad:", age)

	// Short variable declaration := declaration
	secondName, secondAge := "Maria", 22
	fmt.Println("Segundo Nombre:", secondName, "Edad:", secondAge)
}
