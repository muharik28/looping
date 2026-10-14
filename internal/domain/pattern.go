package domain

// Generator abstraction
type Generator interface {
	Generate(size, down, right int) string
}
