package units

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d числовой тип", v)
	case string:
		fmt.Printf("%s string", v)
	case bool:
		fmt.Printf("%t bool", v)
	default:
		fmt.Printf("chan")
	}
}

func EmptyInterface() {
	do(21)
	do("hello")
	do(true)
}
