package units

import (
	"fmt"
	"log"
	"os"
	"time"
)

func WriteCh() {
	var workTime int
	fmt.Print("Введите врямя работы приложения в секундах: ")
	fmt.Fscan(os.Stdin, &workTime)
	to := time.After(time.Duration(workTime) * time.Second)
	channel := make(chan int64)
	go func() {
		for {
			readCh(channel)
		}
	}()
	for {
		select {
		case <-to:
			log.Println("CANCEL")
			return
		default:
			channel <- time.Now().UnixMicro()
			//time.Sleep(1 * time.Millisecond)
		}
	}

}

func readCh(channel chan int64) {
	data := <-channel
	fmt.Println(data)
}
