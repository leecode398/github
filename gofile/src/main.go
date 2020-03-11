package main

import (
	"channel"
	"code"
	"fmt"
	_ "mycontext"
	_ "mysql"
	"pipe"
	_ "workerpool"
	"myhttp"
)

func main() {
	fmt.Println("hello world!")
	code.Test1()
	f := code.Add()
	fmt.Println(f(1))
	fmt.Println(f(2))

	//channel
	channel.Chan1()
	//pipe
	pipe.Pipe()
	//task
	channel.Task()
	//Fibonacci
	//code.Fibonacci()
	//workerpool.Worker()
	//mysql
	//mysql.Mysql()
	//context
	//mycontext.MyContext()

	channel.Chan_test()

	myhttp.Myhttp()
}
