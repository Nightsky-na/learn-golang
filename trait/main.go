package main

import (
	"fmt"

	"errors"
	// "github.com/Nightsky-na/trait/example"
)

type volumer interface {
	Volume() float64
}

type cube struct {
	edge float64
} //edge x edge x edge

type triangularPrism struct {
	base     float64
	attitude float64
	height   float64
} // 0.5 x base x attitude x height

func VolumeOf(v volumer) float64 {
	return v.Volume()
}

func (c cube) Volume() float64 {
	return c.edge * c.edge * c.edge
}

func (t triangularPrism) Volume() float64 {
	return 0.5 * t.base * t.attitude * t.height
}

// ==================================================================
func validateLength(s string) error {
	if len([]rune(s)) < 8 {
		return errors.New("too short string")
	}
	return nil
}

var ageError = errors.New("your age is negative!")

type ErrTooYoung int

func (err ErrTooYoung) Error() string {
	return fmt.Sprintf("%d is too young", err)
}

func validateAge(n int) error {
	if n < 0 {
		return ageError
	}
	if n < 18 {
		return ErrTooYoung(n)
	}
	return nil
}

func main() {
	var a interface{}

	// %T => type
	//

	a = "ten"
	fmt.Printf("%T %v\n", a, a)

	a = true
	fmt.Printf("%T %v\n", a, a)

	a = 10
	fmt.Printf("%T %v\n", a, a)
	// a = func() string { return "hello" }
	// fmt.Printf("%T %v\n", a, a)

	// fmt.Println()
	// example.Execute()

	// var i int
	if i, ok := a.(float32); ok {
		fmt.Println(i)
	} else {
		fmt.Println("Worng type")
	}

	var c, b Obj
	c = rectangle{
		w: 4,
		l: 4,
	}
	b = triangle{
		w: 4,
		h: 4,
	}

	fmt.Println(c.Area())
	fmt.Println(b.Area())

	if v, ok := b.(triangle); ok {
		v.h = 5
	}
}

type Obj interface {
	Area() float64
}

type rectangle struct {
	w, l float64
}

func (rec rectangle) Area() float64 {
	return rec.l * rec.w
}

type triangle struct {
	w, h float64
}

func (tri triangle) Area() float64 {
	return tri.h * tri.w * 0.5
}
