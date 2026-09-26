package main

import "testing"

func TestSquarePatternNoTrailingNewline(t *testing.T) {
	got := SquarePattern{}.Generate(3)
	want := "***\n***\n***"

	if got != want {
		t.Fatalf("Generate() = %q, want %q", got, want)
	}
}

func TestAscendingTrianglePatternNoTrailingNewline(t *testing.T) {
	got := AscendingTrianglePattern{}.Generate(3)
	want := "*\n**\n***"

	if got != want {
		t.Fatalf("Generate() = %q, want %q", got, want)
	}
}

func TestDescendingTrianglePatternNoTrailingNewline(t *testing.T) {
	got := DescendingTrianglePattern{}.Generate(3)
	want := "***\n**\n*"

	if got != want {
		t.Fatalf("Generate() = %q, want %q", got, want)
	}
}

/*
	Result for 5 stars
*/
//   *    *    *    *    *
//  ***  ***  ***  ***  ***
// *************************
//  ***  ***  ***  ***  ***
//   *    *    *    *    *
func TestDiamondPattern(t *testing.T) {
	got := DiamondPattern{}.Generate(5)
	want := "  *    *    *    *    *  \n ***  ***  ***  ***  *** \n*************************\n ***  ***  ***  ***  *** \n  *    *    *    *    *  "

	if got != want {
		t.Fatalf("Generate() = %q, want %q", got, want)
	}
}
