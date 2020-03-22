package main

import (
	"fmt"
	"reflect"
)

const (
	a = iota
	b
	c
	d
	e = 11
	k
	g = iota
	p
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	fmt.Println(a, b, c, d, e, k, g, p)
}
