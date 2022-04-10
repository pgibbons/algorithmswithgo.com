package module01

import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	// return empty case
	if len(word) == 0 {
		return ""
	}

	// reverse order building
	var result_builder strings.Builder
	for i := (len(word) - 1); i >= 0; i-- {
		result_builder.WriteByte(word[i])
	}

	return result_builder.String()
}
