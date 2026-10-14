package application

import (
	"testing"
)

type MockGenerator struct {
	size   int
	down   int
	right  int
	result string
}

func (m *MockGenerator) Generate(size, down, right int) string {

	m.size = size
	m.down = down
	m.right = right

	return m.result
}

func TestPatternGenerator_GeneratePattern(t *testing.T) {
	mockGenerator := &MockGenerator{
		result: "test",
	}

	patternGenerator := NewPatternGenerator(mockGenerator)

	got := patternGenerator.GeneratePattern(5, 2, 3)

	if got != "test" {
		t.Errorf("got %q, want %q", got, "test")
	}

	if mockGenerator.size != 5 {
		t.Errorf("size = %d, want %d", mockGenerator.size, 0)
	}

	if mockGenerator.down != 2 {
		t.Errorf("down = %d, want %d", mockGenerator.down, 0)
	}

	if mockGenerator.right != 3 {
		t.Errorf("right = %d, want %d", mockGenerator.right, 0)
	}
}
