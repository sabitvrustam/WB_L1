package units

import "fmt"

func Revers() {
	s := "Hello world"
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	fmt.Println(result)
}
