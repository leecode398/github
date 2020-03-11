package workerpool

import (
	"fmt"
)

const (
	NUMBER = 10
)

type task struct {
	begin int
	end int
	result chan<- int
}

//task:计算begin到end的和.
//结果写入chan result.
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

func Worker() {
	workers := NUMBER
	//工作通道.
	taskchan := make(chan task, 10)
	//结果通道.
	resultchan := make(chan int, 10)
	//worker信号通道
	done := make(chan struct{}, 10)
	//初始化task的goroutine,计算100个自然数之和.
	go InitTask(taskchan, resultchan, 100)
	//分发任务到NUMBER个goroutine池.
	DistributeTask(taskchan, workers, done)
	//获取各个goroutine处理完任务的通知,并关闭结果通道.
	go CloseResult(done, resultchan, workers)
	//通过结通道获取结果并汇总.
	sum := ProcessResult(resultchan)
	fmt.Println("sum=", sum)
}

func InitTask(taskchan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin : b,
			end : e,
			result : r,
		}
		taskchan <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin : high + 1,
			end : p,
			result : r,
		}
		taskchan <- tsk
	}
	close(taskchan)
}

func DistributeTask(taskchan <-chan task, workers int, done chan struct{}) {
	for i := 0; i < workers; i++ {
		go ProcessTask(taskchan, done)
	}
}

func ProcessTask(taskchan <-chan task, done chan struct{}) {
	for t := range taskchan {
		t.do()
	}
	done <- struct{}{}
}

func CloseResult(done chan struct{}, resultchan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(resultchan)
}

func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
}