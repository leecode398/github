package main

import "fmt"

var sum = func(a, b int) int {
    return a + b
}

func doinput(f func(int, int) int, a, b int) int {
    return f(a, b)
}

func wrap(op string) func(int, int) int {
    switch op {
    case "add":
        return func(a, b int) int {
            return a + b
        }
    case "sub":
        return func(a, b int) int {
            return a - b
        }
    default:
        return nil
    }
}

func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    sum(1, 2)

    doinput(func(x, y int) int {
        return x + y
    },1, 2)

    opFunc := wrap("add")
    re := opFunc(2, 3)
    fmt.Println("%d\n", re)
}
