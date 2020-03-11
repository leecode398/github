package main

import (
	"fmt"
	"sync"
	"time"
    "runtime"
    "os"
    "os/signal"
    "syscall"
)

//用close函数关闭通道引发结束通知,一次性事件
func example1() {
	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Println(id, ":ready.")     //准备就绪
			<-ready                        //等待发令
			fmt.Println(id, ":running...") //开跑
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Ready?Go!")

	close(ready) //传递开始信号

	wg.Wait()
}

//连续或多样性事件可以传递不同数据标志实现,还可以使用sync.Cond实现单播或广播事件
func example2() {
	c := make(chan int, 3)
	c <- 10
	c <- 20
	close(c)

	for i := 0; i < cap(c)+1; i++ {
		x, ok := <-c
		fmt.Println(i, ":", ok, x)
	}
}

//单向,限制收发操作
func example3() {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c
	go func() {
		defer wg.Done()

		for x := range recv {
			fmt.Println(x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()
	wg.Wait()
}

//使用select语句同时处理多个通道
func example4() {
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	go func() {
		defer wg.Done()

		for {
			var (
				name string
				x    int
				ok   bool
			)
			select {
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}
			if !ok {
				return
			}
			fmt.Println(name, x)
		}
	}()
	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			case b <- i * 10:
			}
		}
	}()
	wg.Wait()
}

//等全部通道消息处理结束(closed),可以将已完成通道设置为nil,这样它就阻塞,不会被select选中
func example5() {
	var wg sync.WaitGroup
	wg.Add(3)

	a, b := make(chan int), make(chan int)
	go func() {
		defer wg.Done()

		for {
			select {
			case x, ok := <-a:
				if !ok {
					a = nil
					break
				}
				fmt.Println("a", x)
			case x, ok := <-b:
				if !ok {
					b = nil
					break
				}
				fmt.Println("b", x)
			}
			if a == nil && b == nil {
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		defer close(a)

		for i := 0; i < 3; i++ {
			a <- i
		}
	}()

	go func() {
		defer wg.Done()
		defer close(b)

		for i := 0; i < 5; i++ {
			b <- i * 10
		}
	}()

	wg.Wait()
}

//同一通道也会随机选择case执行
func example6() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)

	go func() {
		defer wg.Done()

		for {
			var v int
			var ok bool

			select {
			case v, ok = <-c:
				fmt.Println("a1", v)
			case v, ok = <-c:
				fmt.Println("a2", v)
			}

			if !ok {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 10; i++ {
			select {
			case c <- i:
			case c <- i * 10:
			}
		}
	}()

	wg.Wait()
}

// 当所有通道都不可用时,select会执行default语句,如此可以避开select阻塞,但须注意处理外层循环,以免陷入空耗.
func example7() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			select {
			case x, ok := <-c:
				if !ok {
					return
				}
				fmt.Println("data:", x)
			default:
			}
			fmt.Println(time.Now())
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 5)

	c <- 100
	close(c)

	<-done
}

//default处理默认罗技
func example8() {
	done := make(chan struct{})

	data := []chan int{ //数据缓冲区
		make(chan int, 3),
	}

	go func() {
		defer close(done)

		for i := 0; i < 10; i++ {
			select {
			case data[len(data)-1] <- i: //生产数据
			default: //当前通道已满,生成新的缓存通道
				data = append(data, make(chan int, 3))
			}
		}
	}()
	<-done

	for i := 0; i < len(data); i++ { //显示所有数据
		c := data[i]
		close(c)

		for x := range c {
			fmt.Println(x)
		}
	}
}

//工厂方法将goroutine和通道绑定
type receiver struct {
	sync.WaitGroup
	data chan int
}

func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}

	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.data {
			fmt.Println("recv:", x)
		}
	}()

	return r
}

func example9() {
	r := newReceiver()
	r.data <- 1
	r.data <- 2

	close(r.data)
	r.Wait()
}

//ID generator,Pool
type pool chan[]byte

func newPool(cap int)pool{
    return make(chan []byte,cap)
}

func (p pool) get() []byte {
    var v[]byte
    select{
    case v= <-p:
    default:
        v = make([]byte,10)
    }
    return v
}

func (p pool) put(b[]byte) {
    select{
    case p<-b:
    default:
    }
}

//通道实现型号量(semaphore)
func example10() {
    fmt.Println(runtime.NumCPU()) // 默认CPU核心数
    runtime.GOMAXPROCS(4)
    var wg sync.WaitGroup

    sem:=make(chan struct{},2)

    for i:=0; i<5; i++ {
        wg.Add(1)

        go func(id int) {
            defer wg.Done()

            sem<-struct{}{}
            defer func() {
                <-sem
            }()

            time.Sleep(time.Second*2)
            fmt.Println(id,time.Now())
        }(i)
    }

    wg.Wait()
}

//标准库time提供tieout和tick channel实现
func example11() {
    go func() {
        for{
            select{
            case <-time.After(time.Second*5):
                fmt.Println("timeout...")
                os.Exit(0)
            }
        }
    }()

    go func() {
        tick:=time.Tick(time.Second)

        for{
            select{
            case <-tick:
                fmt.Println(time.Now())
            }
        }
    }()

    <-(chan struct{})(nil)
}

//捕获INT,TERM信号,实现简易的atexit函数
var exits= &struct{ 
   sync.RWMutex
   funcs  []func() 
   signals chan os.Signal
}{} 
  
func atexit(f func()) { 
   exits.Lock() 
   defer exits.Unlock() 
   exits.funcs=append(exits.funcs,f) 
} 
  
func waitExit() { 
   if exits.signals==nil{ 
       exits.signals=make(chan os.Signal) 
       signal.Notify(exits.signals,syscall.SIGINT,syscall.SIGTERM) 
    } 
  
   exits.RLock() 
   for _,f:=range exits.funcs{ 
       defer f()     // 即便某些函数panic，延迟调用也能确保后续函数执行 
    }              // 延迟调用按FILO顺序执行 
   exits.RUnlock() 
  
    <-exits.signals
} 
  
func example12() { 
   atexit(func() {println("exit1...") }) 
   atexit(func() {println("exit2...") }) 
  
   waitExit() 
}

//goroutine leak
func test() {
    c:=make(chan int)

    for i:=0; i<10; i++ {
        go func() {
            <-c
        }()
    }
}

func example13() {
    test()

    for{
        time.Sleep(time.Second)
        runtime.GC()
    }
}
func main() {
	//example1()
	//example2()
	//example3()
	//example4()
	//example5()
	//example6()
	//example7()
	//example8()
    //example9()
    // example10()
    // example11()
    // example12()
    example13()
}
