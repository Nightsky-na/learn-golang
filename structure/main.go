package main

import "fmt"

type rectangle struct {
	w, l float64
}

// function
func Area(rec rectangle) float64 {
	fmt.Print("Function : ")
	return rec.l * rec.w
}

// ~method
func (rec rectangle) Area() float64 {
	fmt.Print("Similar method : ")
	return rec.l * rec.w
}

type Book struct {
	Name   string
	Author string
}

func (book Book) String() string {
	return fmt.Sprintf("%s by %s", book.Name, book.Author)
}

func (book *Book) SetName(name string) {
	book.Name = name
}

func (rec *rectangle) SetWidth(w float64)  {
	rec.w = w
}

func main() {
	rec := rectangle{
		w: 4,
		l: 5,
	}

	rec.w = 6

	fmt.Println(rec.l)
	fmt.Println(rec.w)

	fmt.Println((Area(rec)))
	fmt.Println((rec.Area()))
	rec.SetWidth(30)
	fmt.Println((rec.Area()))

}
