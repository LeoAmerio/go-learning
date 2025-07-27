package book

import "fmt"

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	title  string
	author string
	pages  int
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

// Composicion
type TextBook struct {
	Book
	editorial string
	niveles   string
}

func NewTextBook(title string, author string, pages int, editorial string, niveles string) *TextBook {
	return &TextBook{
		Book: Book{
			title:  title,
			author: author,
			pages:  pages,
		},
		editorial: editorial,
		niveles:   niveles,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nNiveles: %s\n", b.title, b.author, b.pages, b.editorial, b.niveles)
}
