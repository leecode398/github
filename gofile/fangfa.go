package main

import "fmt"
type SliceInt []int

func (s SliceInt) Sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func SliceInt_Sum(s SliceInt) int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}
type T struct{
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(i int) {
	t.a = i
}

func (t *T) Print() {
	fmt.Println(t.a)
}

func main() {
	var s SliceInt = []int{1,2,3,4}
	s.Sum()
	fmt.Println(SliceInt_Sum(s))
	fmt.Println(s.Sum())

	var t = T{}
	//t.Set(2)
	(*T).Set(&t,4)
	fmt.Println(t.Get())
	t.Print()
}