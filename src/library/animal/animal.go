package animal

import "fmt"

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre, " says Guau Guau")
}

type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre, " says Miau Miau")
}

func HacerRuido(animal Animal) {
	animal.Sonido()
}
