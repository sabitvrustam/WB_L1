package units

import "fmt"

func CreateSet() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)
	for _, v := range input {
		set[v] = true
	}

	fmt.Println(set)
}
