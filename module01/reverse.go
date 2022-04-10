package module01

import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	// reverse order building
	var result_builder strings.Builder
	runes := []rune(word)
	for i := (len(runes) - 1); i >= 0; i-- {
		result_builder.WriteRune(runes[i])
	}
	return result_builder.String()
}
