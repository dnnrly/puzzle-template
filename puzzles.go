package puzzle

// Solution generates the solution to a single puzzle, or part of one
type Solution func() int

// Puzzle is a function that can be called to produce a Euler puzzle solution
type Puzzle struct {
	Parts []Solution
}

// Size gives the number of parts for this puzzle
func (p *Puzzle) Size() int {
	return len(p.Parts)
}

// RegisterPuzzles is the list of all puzzles with solutions
func RegisterPuzzles() []Puzzle {
	puzzles := []Puzzle{}

	// next puzzle

	return puzzles
}
