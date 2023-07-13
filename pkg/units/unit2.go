package units

import (
	"fmt"
	"math"
	"sync"
)

func SquaresValue() {
	channel := make(chan float64)
	wg := sync.WaitGroup{}
	nums := [...]int{2, 4, 6, 8, 10}
	wg.Add(len(nums))
	for _, num := range nums {
		num := num
		go func() {
			square(num, channel)
			wg.Done()
		}()
	}
	close := make(chan interface{})
	go func() {
		for {
			select {
			case res := <-channel:
				fmt.Println(res)
			case <-close:
				return
			}
		}
	}()
	wg.Wait()
	close <- struct{}{}
}

func square(num int, channel chan float64) {
	result := math.Pow(float64(num), 2.0)
	channel <- result
}
