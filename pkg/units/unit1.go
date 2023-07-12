package units

import "time"

type Human struct {
	Name string
	Age  time.Time
}

type Action struct {
	Human
}

func (h *Human) GetAge() int {
	return time.Now().Day() - h.Age.Day()
}
