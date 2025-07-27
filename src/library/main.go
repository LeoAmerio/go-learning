package main

import (
	"library/animal"
)

func main() {
	// myBook := book.NewBook("The Go Programming Language", "Alan A. A. Donovan and Brian W. Kernighan", 400)
	// myBook.SetTitle("The Go Programming Language - Updated Edition")

	// var secundBook = book.Book{
	// 	title:  "The Go Programming Language",
	// 	author: "Alan A. A. Donovan and Brian W. Kernighan",
	// 	pages:  400,
	// }
	// secundBook.PrintInfo()

	// fmt.Println("AAA", myBook.GetTitle())
	// myBook.PrintInfo()

	// myTextBook := book.NewTextBook("Introduction to Algorithms", "Thomas H. Cormen", 1312, "MIT Press", "Undergraduate")
	// // myTextBook.PrintInfo()

	// book.Print(myBook)
	// book.Print(myTextBook)

	// miPerro := animal.Perro{ Nombre: "Rex" }
	// miGato := animal.Gato{ Nombre: "Michi" }
	// animal.HacerRuido(&miPerro)
	// animal.HacerRuido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Rex"},
		&animal.Gato{Nombre: "Michi"},
		&animal.Perro{Nombre: "Firulais"},
		&animal.Gato{Nombre: "Luna"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
