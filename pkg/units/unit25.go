package units

import (
	"fmt"
	"time"
)

func CallSlip() {
	fmt.Println("начало Slip ожидание 10 секунд")
	slip(5 * time.Second)
	fmt.Println("конец Slip")
}

func slip(d time.Duration) {
	select {
	case <-time.After(d):
	}
}
