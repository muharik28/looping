package application

import "github.com/muharik28/looping/internal/domain"

type PatternGenerator struct {
	generator domain.Generator
}

func NewPatternGenerator(generator domain.Generator) *PatternGenerator {

	return &PatternGenerator{
		generator: generator,
	}
}

func (p *PatternGenerator) GeneratePattern(size, down, right int) string {

	return p.generator.Generate(size, down, right)
}
