package pattern_generator

import (
	"fmt"
	"os"

	"github.com/muharik28/looping/input"
	"github.com/muharik28/looping/internal/application"
	"github.com/muharik28/looping/internal/infrastructure/patterns"
)

func Execute() {
	console := input.NewConsole()

	size, err := console.ReadInt("Enter size of pattern: ")
	if err != nil {
		fmt.Println("Invalid size.")
		os.Exit(1)
	}

	down, err := console.ReadInt("Enter down of pattern: ")
	if err != nil {
		fmt.Println("Invalid down.")
		os.Exit(1)
	}

	right, err := console.ReadInt("Enter right of pattern: ")
	if err != nil {
		fmt.Println("Invalid right.")
		os.Exit(1)
	}

	var pyramidPattern patterns.PyramidPattern

	generator := application.NewPatternGenerator(&pyramidPattern)

	result := generator.GeneratePattern(size, down, right)

	fmt.Print(result)
}
