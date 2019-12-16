package main

import "fmt"

/*
继承
一个结构体嵌到另一个结构体，称作组合
匿名和组合的区别
如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现继承
如果一个struct嵌套了另一个【有名】的结构体，那么这个模式叫做组合
如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现多重继承
*/

type Car struct {
    weight int
    name   string
}

func (p *Car) Run() {
    fmt.Println("running")
}

type Bike struct {
    Car
    lunzi int
}
type Train struct {
    Car
}

func (p *Train) String() string {
    str := fmt.Sprintf("name=[%s] weight=[%d]", p.name, p.weight)
    return str
}

func main() {
    var a Bike
    a.weight = 100
    a.name = "bike"
    a.lunzi = 2
    fmt.Println(a)
    a.Run()

    var b Train
    b.weight = 100
    b.name = "train"
    b.Run()
    fmt.Printf("%s", &b)
}