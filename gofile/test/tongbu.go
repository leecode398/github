package main

import(
	"fmt"
	"sync"
	"time"
)

type data struct{
	sync.Mutex
}

func (d data)test(s string) {
	d.Lock()
	defer d.Unlock()

	for i:=0; i<5; i++ {
		fmt.Println(s,i)
		time.Sleep(time.Second)
	}
}

func exam1() {
	var wg sync.WaitGroup
	wg.Add(2)

	var d data

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()
	wg.Wait()
}

func main() {
	exam1()
}