package puzzle_test

import (
	"fmt"
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

func ExamplePuzzle_initialisation() {
	NewPuzzleExample := func() puzzle.Puzzle {
		data := []int{
			100, 101, 101,
		}

		return puzzle.Puzzle{
			Parts: []puzzle.Solution{
				func() int { return data[0] },
				func() int { return data[1] },
				func() int { return data[2] },
			},
		}
	}

	p := NewPuzzleExample()

	for i, s := range p.Parts {
		fmt.Printf("Solution %d is %d\n", i, s())
	}

	// Output:
	// Solution 0 is 100
	// Solution 1 is 101
	// Solution 2 is 101
}
