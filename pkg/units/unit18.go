package units

import (
	"log"
	"sync"
)

type Count struct {
	num int
	sync.Mutex
}

func (c *Count) Inc() {
	c.Lock()
	defer c.Unlock()

	c.num += 1
}

func (c *Count) Value() int {
	return c.num
}

func Counter() {
	cnt := &Count{
		num: 0,
	}

	finish := make(chan struct{})

	go Do(cnt, finish)

	select {
	case <-finish:
		log.Printf("показания счетчика: %d", cnt.Value())
	}
}

func Do(cnt *Count, finish chan struct{}) {
	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(num int, cnt *Count, wg *sync.WaitGroup) {
			defer wg.Done()
			cnt.Inc()
		}(i, cnt, &wg)
	}

	wg.Wait()
	finish <- struct{}{}
	close(finish)
}
