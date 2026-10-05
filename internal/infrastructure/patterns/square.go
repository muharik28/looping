package patterns

import "strings"

type SquarePattern struct{}

func (s *SquarePattern) Generate(size, down, right int) string {

	if size <= 0 || down <= 0 || right <= 0 {
		return ""
	}

	var result strings.Builder

	line := strings.Repeat("*", size*right)

	for d := 0; d < down; d++ {
		result.WriteString(line)

		if d < down-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}
