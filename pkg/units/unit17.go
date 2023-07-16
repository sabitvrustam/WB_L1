package units

import (
	"fmt"
	"sort"
)

func BinarySearch() {

	a := []string{"A", "C", "C"}
	x := "C"

	i := sort.Search(len(a), func(i int) bool { return x <= a[i] })
	if i < len(a) && a[i] == x {
		fmt.Printf("Найдено %s по индексу %d в %v.\n", x, i, a)
	} else {
		fmt.Printf("Не найдено %s в %v.\n", x, a)
	}
}
