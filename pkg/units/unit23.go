package units

import (
	"fmt"
)

func RemoveElementSlice() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	copy(a[i:], a[i+1:])

	a[len(a)-1] = ""

	a = a[:len(a)-1]

	fmt.Println(a) // [A B D E]
}
