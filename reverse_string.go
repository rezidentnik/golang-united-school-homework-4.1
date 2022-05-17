package reverse_string

import (
	"unicode/utf8"
)

func ReverseString(input string) (output string) {
	runesNumber := utf8.RuneCountInString(input)
	reversedSlice := make([]rune, runesNumber)
	i := runesNumber

	for _, rune := range input {
		i--
		reversedSlice[i] = rune
	}

	output = string(reversedSlice)

	return output
}
