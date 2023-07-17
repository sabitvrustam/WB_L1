package units

import "fmt"

func SetIntersection() {
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	set2 := map[int]bool{1: true, 2: true, 3: true, 40: true, 50: true}

	intersection := make(map[int]bool)
	for k := range set1 {
		if _, found := set2[k]; found {
			intersection[k] = true
		}
	}

	fmt.Println(intersection)
}
