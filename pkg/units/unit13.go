package units

import "fmt"

func ChangeVar() {
	a := 2
	b := 10

	fmt.Printf("Дано два числа %d и %d замена местами без дополнительной переменной", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf(" результат %d и %d ", a, b)
}
