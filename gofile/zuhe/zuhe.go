package main

import "fmt"

type X struct {
	a int
}

type Y struct {
	X
	b int
}

type Z struct {
	Y
	c int
}

func (x X) Print() {
	fmt.Printf("Ln X, a =%d\n", x.a)
}

func (x X) XPrint() {
	fmt.Printf("In X, a =%d\n", x.a)
}

func (y Y) YPrint() {
	fmt.Printf("In Y, a =%d\n", y.a)
}

func (z Z) ZPrint() {
	fmt.Printf("In Z, a =%d\n", z.a)
	z.Y.Print()
	z.Y.X.Print()
}
func main() {
	x := X{a: 1}

	y := Y{
		X: x,
		b: 2,
	}

	z := Z{
		Y: y,
		c : 3,
	}

	z.Print()
	z.XPrint()
	z.Y.XPrint()

}