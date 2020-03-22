package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 0}

	for m := range b {
		fmt.Print(m)
	}

	for _, v := range a {
		fmt.Print(v, " ")
		for _, m := range b {
			fmt.Print(m, " ")
			if m == 7 {
				break
			}
		}
		fmt.Println("")
	}
}
