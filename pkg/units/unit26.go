package units

import (
	"log"
	"strings"
)

func CountLetters(str string) (bool, string) {
	str = strings.ToLower(str)
	slcstr := strings.Split(str, "")
	mpbool := make(map[string]bool, 0)
	for _, item := range slcstr {
		_, ok := mpbool[item]
		if ok {
			return false, str
		}
		mpbool[item] = true
	}
	return true, str
}

func UniqueSymbols() {
	log.Println(CountLetters("akd,i"))
}
