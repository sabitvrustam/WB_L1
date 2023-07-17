package units

import "fmt"

type MessageSender struct {
}

func (m *MessageSender) SendMsg(msg string) {
	fmt.Println()
}

type Adapter struct {
	prefix string
}

func NewAdapter(prefix string) *Adapter {
	return &Adapter{prefix: prefix}
}

func (a *Adapter) SendMsg(msg string) {
	fmt.Println(a.prefix + msg)
}

func Unit21() {
	adapter := NewAdapter("test_prefix")
	adapter.SendMsg("message from adapter")
}
