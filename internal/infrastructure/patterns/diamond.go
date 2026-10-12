package patterns

import "strings"

type DiamondPattern struct{}

func (d DiamondPattern) Generate(size, down, right int) string {

	if size <= 0 || down <= 0 || right <= 0 {
		return ""
	}

	var rows []string

	// size*3 harcode 3 for result stars 3 if 1 stars change 3 to 1, this is row
	for x := 0; x < size*down; x++ {
		var line strings.Builder

		// size*5 harcode 5 for result stars 5 if 1 stars change 5 to 1, this is column
		for y := 0; y < size*right; y++ {
			line.WriteString(d.diamondCharacter(x, y, size))
		}
		rows = append(rows, line.String())
	}

	result := strings.Join(rows, "\n")

	return result
}

func (DiamondPattern) diamondCharacter(x, y, size int) string {

	mid := size / 2

	xOffset := x % size
	yOffset := y % size

	distanceX := xOffset - mid
	distanceY := yOffset - mid

	if distanceX < 0 {
		distanceX = -distanceX
	}

	if distanceY < 0 {
		distanceY = -distanceY
	}

	if distanceX+distanceY <= mid {
		return "*"
	}

	return " "
}
