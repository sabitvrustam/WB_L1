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

func CloseGoroutine() {
	var closemetod int
	fmt.Println("Введите способ закрытия Goroutine")
	fmt.Println("1 - закрытие с помощю канала")
	fmt.Println("2 - закрытие с Context")
	fmt.Fscan(os.Stdin, &closemetod)
	switch closemetod {
	case 1:
		metodCannal()
	case 2:
		metodContext()
	default:
		fmt.Println("не выбран не один из вариантов")
		return
	}
}

func metodCannal() {
	fmt.Println("metod channal")
	var workTime int
	fmt.Print("Введите врямя работы приложения в секундах: ")
	fmt.Fscan(os.Stdin, &workTime)
	to := time.After(time.Duration(workTime) * time.Second)
	channel := make(chan int64)
	go readCh(channel)

	for {
		select {
		case <-to:
			log.Println("CANCEL")
			break
		default:
			channel <- time.Now().UnixMicro()
			time.Sleep(1 * time.Second)
		}
	}
}

func metodContext() {
	fmt.Println("metod context, for exit ctr-c")
	channel := make(chan int64)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)

	for id := 0; id <= 5; id++ {
		go func(id int) {
			worker(id, channel, ctx)
		}(id)
	}

	for {
		select {
		case <-ctx.Done():
			cancel()
			log.Println("CANCEL")
			break
		default:
			channel <- time.Now().UnixMicro()
			time.Sleep(time.Second)
		}
	}
}
