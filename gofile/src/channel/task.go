package channel

import (
	"fmt"
	"sync"
)

type task struct {
	begin	int
	end		int
	result	chan<- int
}

func (t *task) do(){
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}	

func Task() {
	taskchan := make(chan task, 10)
	resultchan := make(chan int, 10)
	
	wait := &sync.WaitGroup{}

	go InitTask(taskchan, resultchan, 100)

	go DistributeTask(taskchan, wait, resultchan)

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

func DistributeTask(taskchan <-chan task, wait *sync.WaitGroup, result chan int) {
	for v := range taskchan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	close(result)
}

func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
}

