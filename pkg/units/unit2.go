package units

import (
	"fmt"
	"math"
	"sync"
)

func SquaresValue() {
	wg := sync.WaitGroup{}
	nums := [...]int{2, 4, 6, 8, 10}
	wg.Add(len(nums))
	for _, num := range nums {
		num := num
		go func() {
			fmt.Println(square(num))
			wg.Done()
		}()
	}
	wg.Wait()
}

func square(num int) float64 {
	return math.Pow(float64(num), 2.0)
}
