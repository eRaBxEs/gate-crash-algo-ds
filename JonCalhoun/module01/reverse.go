package module01

import "fmt"

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	var newSt string
	for _, st := range word {
		newSt = fmt.Sprintf("%s%s", string(st), newSt)
		//newSt = string(st) + newSt
	}

	return newSt
}
