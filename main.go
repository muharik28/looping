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
	fmt.Print("Enter size of square pattern: ")

	var size string
	fmt.Scanln(&size)

	i, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	squarePattern := SquarePattern{}

	generator := NewPatternGenerator(squarePattern)

	line := generator.GeneratePattern(i)
	fmt.Print(line)
}
