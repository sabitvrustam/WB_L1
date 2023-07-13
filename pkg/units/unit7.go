package units

import (
	"fmt"
	"sync"
)

func WrieDataMap() {
	c := map[int]int{}
	wg := sync.WaitGroup{}
	n := 200
	m := sync.Mutex{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			m.Lock()
			c[i] = i
			m.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(len(c))
}
