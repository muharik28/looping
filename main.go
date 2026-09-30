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
	Generate(size int) string
}

// SquarePattern implementation
type SquarePattern struct{}

func (SquarePattern) Generate(size int) string {
	line := strings.Repeat("*", size)
	rows := make([]string, size)
	for i := range size {
		rows[i] = line
	}

	result := strings.Join(rows, "\n")

	return result
}

type AscendingTrianglePattern struct{}

func (AscendingTrianglePattern) Generate(size int) string {
	var line string
	rows := make([]string, size)
	for i := range size {
		line += "*"
		rows[i] = line
	}

	result := strings.Join(rows, "\n")

	return result
}

type DescendingTrianglePattern struct{}

func (DescendingTrianglePattern) Generate(size int) string {
	line := strings.Repeat("*", size)
	rows := make([]string, size)
	for i := range size {
		rows[i] = line
		line = line[:len(line)-1]
	}

	result := strings.Join(rows, "\n")

	return result
}

type DiamondPattern struct{}

func (d DiamondPattern) Generate(size int) string {
	var rows []string

	// size*3 harcode 3 for result stars 3 if 1 stars change 3 to 1, this is row
	for x := 0; x < size*3; x++ {
		var line strings.Builder

		// size*5 harcode 5 for result stars 5 if 1 stars change 5 to 1, this is column
		for y := 0; y < size*5; y++ {
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

func (p *PatternGenerator) GeneratePattern(size int) string {
	return p.generator.Generate(size)
}

func main() {
	fmt.Print("Enter size of pattern: ")

	var size string
	fmt.Scanln(&size)

	i, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	diamondPattern := DiamondPattern{}

	generator := NewPatternGenerator(diamondPattern)

	line := generator.GeneratePattern(i)
	fmt.Print(line)
}
