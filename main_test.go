package main

import "testing"

/*
	Result for 3 square pattern
*/
// ***
// ***
// ***
func TestSquarePatternNoTrailingNewline(t *testing.T) {
	got := SquarePattern{}.Generate(3, 3, 1)
	want := "***\n***\n***"

	if got != want {
		t.Fatalf("Generate() = %q, want %q", got, want)
	}
}

/*
	Result for 5 ascending triangle pattern
*/
// *  *  *  *  *
// ** ** ** ** **
// ***************
func TestAscendingTrianglePatternNoTrailingNewline(t *testing.T) {
	got := AscendingTrianglePattern{}.Generate(3, 1, 5)
	want := "*  *  *  *  *\n** ** ** ** **\n***************"

	if got != want {
		t.Fatalf("Generate() = %q, want %q", got, want)
	}
}

/*
	Result for 5 descending triangle pattern
*/
// ***************
// ** ** ** ** **
// *  *  *  *  *
func TestDescendingTrianglePatternNoTrailingNewline(t *testing.T) {
	got := DescendingTrianglePattern{}.Generate(3, 1, 5)
	want := "***************\n** ** ** ** **\n*  *  *  *  *"

	if got != want {
		t.Fatalf("Generate() = %q, want %q", got, want)
	}
}

/*
	Result for 5 diamond pattern
*/
//   *    *    *    *    *
//  ***  ***  ***  ***  ***
// *************************
//  ***  ***  ***  ***  ***
//   *    *    *    *    *
func TestDiamondPattern(t *testing.T) {
	got := DiamondPattern{}.Generate(5, 1, 5)
	want := "  *    *    *    *    *  \n ***  ***  ***  ***  *** \n*************************\n ***  ***  ***  ***  *** \n  *    *    *    *    *  "

	if got != want {
		t.Fatalf("Generate() = %q, want %q", got, want)
	}
}
