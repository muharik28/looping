package main

import (
	"fmt"
	"strconv"
	"strings"
)

// =============================
// Domain Layer
// =============================

// Generator abstraction
type Generator interface {
	Generate(size, down, right int) string
}

// SquarePattern implementation
type SquarePattern struct{}

func (SquarePattern) Generate(size, down, right int) string {

	if size <= 0 || down <= 0 || right <= 0 {
		return ""
	}

	var result strings.Builder

	line := strings.Repeat("*", size*right)
	for d := range down {

		result.WriteString(line)

		if d < down-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}

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

type DescendingTrianglePattern struct{}

func (DescendingTrianglePattern) Generate(size, down, right int) string {

	if size <= 0 || down <= 0 || right <= 0 {
		return ""
	}

	var result strings.Builder

	for d := range down {

		for row := size; row >= 1; row-- {

			for col := range right {
				result.WriteString(strings.Repeat("*", row))

				if col < right-1 {
					spaces := size - row
					result.WriteString(strings.Repeat(" ", spaces))
				}
			}

			if row > 1 {
				result.WriteString("\n")
			}
		}

		if d < down-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}

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

// =============================
// Application Layer
// =============================

// PatternGenerator use Generator as depedency
type PatternGenerator struct {
	generator Generator
}

// NewPatternGenerator is a constructor for PatternGenerator
//
// Dependency Injection pass through the paremeter generator
func NewPatternGenerator(generator Generator) *PatternGenerator {

	return &PatternGenerator{
		generator: generator,
	}
}

func (p *PatternGenerator) GeneratePattern(size, down, right int) string {

	return p.generator.Generate(size, down, right)
}

func main() {

	fmt.Print("Enter size of pattern: ")

	var size string

	fmt.Scanln(&size)

	s, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	fmt.Print("Enter down of pattern: ")

	var down string

	fmt.Scanln(&down)

	d, err := strconv.Atoi(down)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	fmt.Print("Enter right of pattern: ")

	var right string

	fmt.Scanln(&right)

	r, err := strconv.Atoi(right)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	squarePattern := SquarePattern{}

	generator := NewPatternGenerator(squarePattern)

	line := generator.GeneratePattern(s, d, r)
	fmt.Print(line)
}
