package units

import (
	"fmt"
	"sort"
)

func Sort() {
	s := []int{4, 2, 3, 1, 5, 10}
	sort.Ints(s)
	fmt.Println(s)
}
