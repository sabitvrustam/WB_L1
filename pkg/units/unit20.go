package units

import (
	"fmt"
	"strings"
)

func WordConvection() {
	s := "snow dog sun"
	result := ""

	worlds := strings.Fields(s)
	for _, v := range worlds {
		result = v + " " + result
	}
	fmt.Println(result)

}
