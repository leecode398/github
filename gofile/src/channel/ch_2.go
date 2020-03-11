package channel

import (
	"fmt"
)
func hello(a chan bool) {
	fmt.Println("hello world!!")
	a <- true
}
func Chan_test() {
	a := make(chan bool)
	go hello(a)
	if(<- a) {
		fmt.Println("ok")
	}
	close(a)
}