package reverse_string

import (
	"strings"
)

func ReverseString(input string) (output string) {
	var sb strings.Builder

	for _, rune := range input {
		sb.WriteRune(rune)
	}

	output = sb.String()

	return output
}
