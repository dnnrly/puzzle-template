package puzzle_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dnnrly/puzzle-template"
)

func TestPuzzleHasCorrectCount(t *testing.T) {
	p := puzzle.Puzzle{
		Parts: []puzzle.Solution{
			func() int { return 1 },
			func() int { return 2 },
			func() int { return 3 },
		},
	}

	assert.Equal(t, 3, p.Size())
}

func ExamplePuzzle_create_a_simple_puzzle_solution() {
	// A new puzzle constructor will be generated when you run the 'next' command. Simply
	// fill it in to generate your solution
	NewPuzzleExample := func() puzzle.Puzzle {
		return puzzle.Puzzle{
			Parts: []puzzle.Solution{
				// Add functions here for all of the different puzzle parts for your challenge
				func() int {
					// You can do your calculation here
					result := 10 * 10
					return result
				},
			},
		}
	}

	// Here endeth the example, the tooling will handle plumbing in the constructor
	// automatically.

	p := NewPuzzleExample()

	for i, s := range p.Parts {
		fmt.Printf("Solution %d is %d\n", i, s())
	}

	// Output:
	// Solution 0 is 100
}

func ExamplePuzzle_complex_initialisation_with_helper_functions() {
	NewPuzzleExample := func() puzzle.Puzzle {
		// In this function, you can add large strings (your input data, perhaps) so
		// that they can be parsed in the initialisation phase of the puzzle and won't
		// count to the time taken for your algorithm.
		data := `1
2
3`

		// You can add multiple variables that you use later on if you wish.
		values := []int{}

		// And you can even add helpers that can be called only from inside your
		// puzzle solution.
		convert := func(s string) int {
			v, _ := strconv.Atoi(s)
			return v
		}

		return puzzle.Puzzle{
			Init: func() {
				// You can do the heavy pre-processing here
				lines := strings.Split(data, "\n")
				for _, l := range lines {
					values = append(values, convert(l))
				}
			},
			Parts: []puzzle.Solution{
				// You can consume the pre-processed data in the solutions here
				func() int { return values[0] + 100 },
				func() int { return values[1] + 100 },
				func() int { return values[2] + 100 },
			},
			Tidy: func() {
				// In this function you can clear down any items that you've created
				// in memory. In this example, setting the slice to nil allows the
				// garbage collector to release it. Hopefully this function will allow
				// you to be able to run ALL of your puzzles together.
				values = nil
			},
		}
	}

	p := NewPuzzleExample()

	p.Init()
	for i, s := range p.Parts {
		fmt.Printf("Solution %d is %d\n", i, s())
	}

	// Output:
	// Solution 0 is 101
	// Solution 1 is 102
	// Solution 2 is 103
}
