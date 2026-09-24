package main

import "testing"

func TestSquarePatternNoTrailingNewline(t *testing.T) {
	got := SquarePattern{}.Generate(3)
	want := "***\n***\n***"

	if got != want {
		t.Fatalf("Generate() = %q, want %q", got, want)
	}
}
