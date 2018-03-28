package benchmark_parallel

import (
	"sync"
	"time"
	"sync/atomic"
	"fmt"
)

type worker func(index int) error

func Run(goroutines int, worker worker)  {
	var wg sync.WaitGroup
	ch := make(chan int64, goroutines)
	var ops uint64

	for i:= 0 ; i< goroutines; i++ {
		index := i
		wg.Add(1)
		go func() {
			t0 := time.Now()

			// custom code
			err := worker(index)
			if err != nil {
				atomic.AddUint64(&ops, 1)
			}

			diff := time.Now().Sub(t0)
			ch <- diff.Nanoseconds()

			wg.Done()
		}()
	}

	wg.Wait()

	close(ch)
	var sum int64 = 0
	for nano := range ch {
		sum += nano
	}

	fmt.Printf("%d goroutines, %d / %d error rate, %d nanoseconds/op", goroutines, ops, goroutines, sum/int64(goroutines))
}
