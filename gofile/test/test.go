package main

import (
	"fmt"
)

func findContinuousSequence(target int) [][]int {
    a := [][]int{}
    row := 0
    for left, right := 1, 2
	sum := left + right
    for right < target/2+1 {
        col := 0
        if sum == target {
            for i := left; i <= right; i++ {
                a[row][col] = i
                col++
            }
        } else if sum < terget {
            right++
        } else {
            left++
        }
        row++
    }
    return a
}

func main() {
	target := 15
	fmt.Printf("%v\n",findContinuousSequence(target))
}
