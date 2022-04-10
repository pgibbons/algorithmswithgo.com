package module01

import (
	"fmt"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	var result_builder strings.Builder
	for i := 1; i <= n; i++ {
		if i == 1 {
			result_builder.WriteString("1")
		} else {
			result_builder.WriteString(", " + checkDivisionBy3And5(i))
		}
	}

	fmt.Println(result_builder.String())
}

func checkDivisionBy3And5(num int) string {
	switch {
	case (num%3) == 0 && (num%5) == 0:
		return "Fizz Buzz"
	case (num % 3) == 0:
		return "Fizz"
	case (num % 5) == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", num)
	}
}
