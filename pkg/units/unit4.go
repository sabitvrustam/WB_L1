package units

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WorkerPul() {
	var namberWorker int
	fmt.Print("Введите колличество воркеров: ")
	fmt.Fscan(os.Stdin, &namberWorker)
	channel := make(chan int64)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)

	go func() {
		for id := 0; id <= namberWorker; id++ {
			go func(id int) {
				worker(id, channel, ctx)
			}(id)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			cancel()
			log.Println("CANCEL")
			return
		default:
			channel <- time.Now().UnixMicro()
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker(id int, channel chan int64, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			data, ok := <-channel
			if !ok {
				println(data)
				return
			}

		default:
			data := <-channel
			fmt.Println("worker", id, "Данные", data)
		}
	}
}
