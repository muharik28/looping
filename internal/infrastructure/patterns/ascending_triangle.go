package patterns

import "strings"

type AscendingTrianglePattern struct{}

func (AscendingTrianglePattern) Generate(size, down, right int) string {

	if size <= 0 || down <= 0 || right <= 0 {
		return ""
	}

	var result strings.Builder

	for d := range down {

		for row := 1; row <= size; row++ {

			for col := range right {
				result.WriteString(strings.Repeat("*", row))

				if col < right-1 {
					spaces := size - row
					result.WriteString(strings.Repeat(" ", spaces))
				}
			}

			if row < size {
				result.WriteString("\n")
			}
		}

		if d < down-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}
