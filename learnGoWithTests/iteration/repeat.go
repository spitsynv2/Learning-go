package iteration

import "strings"

// Repeat a given character iterations times
func Repeat(character string, iterartions int) string {
	var result strings.Builder
	for i := 0; i < iterartions; i++ {
		result.WriteString(character)
	}
	return result.String()
}
