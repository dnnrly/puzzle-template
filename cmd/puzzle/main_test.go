package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dnnrly/puzzle-template"
)

func testPuzzles() []puzzle.Puzzle {
	return []puzzle.Puzzle{
		puzzle.Puzzle{
			Init: func() {},
			Parts: []puzzle.Solution{
				func() int { return 1 },
				func() int { return 2 },
				func() int { return 3 },
			},
			Tidy: func() {},
		},
		puzzle.Puzzle{
			Parts: []puzzle.Solution{
				func() int { return 4 },
				func() int { return 5 },
				func() int { return 6 },
			},
			Tidy: func() {},
		},
		puzzle.Puzzle{
			Init: func() {},
			Parts: []puzzle.Solution{
				func() int { return 7 },
				func() int { return 8 },
				func() int { return 9 },
			},
		},
	}
}

func TestRuns0TestsWithtoutError(t *testing.T) {
	logs := []string{}
	config := &Config{
		args: []string{".binary"},
		logf: func(f string, a ...interface{}) {
			t.Logf(f, a...)
			logs = append(logs, fmt.Sprintf(f, a...))
		},
		puzzles: []puzzle.Puzzle{},
	}
	err := Run(config)

	assert.NoError(t, err)
	assert.Contains(t, logs, "There are 0 solutions")
}

func TestRuns3PuzzlesWithtoutError(t *testing.T) {
	logs := []string{}
	config := &Config{
		args: []string{".binary"},
		logf: func(f string, a ...interface{}) {
			t.Logf(f+"\n", a...)
			logs = append(logs, fmt.Sprintf(f, a...))
		},
		puzzles: testPuzzles(),
	}
	err := Run(config)

	assert.NoError(t, err)
	assert.Contains(t, logs, "There are 3 solutions")
	assert.Contains(t, logs, "000   1 Init")
	assert.Contains(t, logs, "000   1.1 - 1")
	assert.Contains(t, logs, "000   1.2 - 2")
	assert.Contains(t, logs, "000   1.3 - 3")
	assert.Contains(t, logs, "000   1 Tidy")
	assert.NotContains(t, logs, "000   2 Init")
	assert.Contains(t, logs, "000   2.1 - 4")
	assert.Contains(t, logs, "000   2.2 - 5")
	assert.Contains(t, logs, "000   2.3 - 6")
	assert.Contains(t, logs, "000   2 Tidy")
	assert.Contains(t, logs, "000   3 Init")
	assert.Contains(t, logs, "000   3.1 - 7")
	assert.Contains(t, logs, "000   3.2 - 8")
	assert.Contains(t, logs, "000   3.3 - 9")
	assert.NotContains(t, logs, "000   3 Tidy")
}

func TestRunsLatestWithtoutError(t *testing.T) {
	logs := []string{}
	config := &Config{
		args: []string{".binary", "latest"},
		logf: func(f string, a ...interface{}) {
			t.Logf(f+"\n", a...)
			logs = append(logs, fmt.Sprintf(f, a...))
		},
		puzzles: testPuzzles(),
	}
	err := Run(config)

	assert.NoError(t, err)
	assert.Contains(t, logs, "There are 3 solutions")
	assert.NotContains(t, logs, "000   1 Init")
	assert.NotContains(t, logs, "000   1.1 - 1")
	assert.NotContains(t, logs, "000   1.2 - 2")
	assert.NotContains(t, logs, "000   1.3 - 3")
	assert.NotContains(t, logs, "000   1 Tidy")
	assert.NotContains(t, logs, "000   2.1 - 4")
	assert.NotContains(t, logs, "000   2.2 - 5")
	assert.NotContains(t, logs, "000   2.3 - 6")
	assert.NotContains(t, logs, "000   2 Tidy")
	assert.Contains(t, logs, "000   3 Init")
	assert.Contains(t, logs, "000   3.1 - 7")
	assert.Contains(t, logs, "000   3.2 - 8")
	assert.Contains(t, logs, "000   3.3 - 9")
}

func TestRunsSpecificTest(t *testing.T) {
	logs := []string{}
	config := &Config{
		args: []string{".binary", "2"},
		logf: func(f string, a ...interface{}) {
			t.Logf(f+"\n", a...)
			logs = append(logs, fmt.Sprintf(f, a...))
		},
		puzzles: testPuzzles(),
	}
	err := Run(config)

	assert.NoError(t, err)
	assert.Contains(t, logs, "There are 3 solutions")
	assert.NotContains(t, logs, "000   1 Init")
	assert.NotContains(t, logs, "000   1.1 - 1")
	assert.NotContains(t, logs, "000   1.2 - 2")
	assert.NotContains(t, logs, "000   1.3 - 3")
	assert.NotContains(t, logs, "000   1 Tidy")
	assert.Contains(t, logs, "000   2.1 - 4")
	assert.Contains(t, logs, "000   2.2 - 5")
	assert.Contains(t, logs, "000   2.3 - 6")
	assert.Contains(t, logs, "000   2 Tidy")
	assert.NotContains(t, logs, "000   3 Init")
	assert.NotContains(t, logs, "000   3.1 - 7")
	assert.NotContains(t, logs, "000   3.2 - 8")
	assert.NotContains(t, logs, "000   3.3 - 9")
}

func TestRunFailsWithUnparsableParameters(t *testing.T) {
	logs := []string{}
	config := &Config{
		args: []string{".binary", "rubbish"},
		logf: func(f string, a ...interface{}) {
			fmt.Printf(f+"\n", a...)
			logs = append(logs, fmt.Sprintf(f, a...))
		},
		puzzles: testPuzzles(),
	}
	err := Run(config)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "'rubbish' is not an option")
}

func TestRunFailsWithUnknownTest(t *testing.T) {
	logs := []string{}
	config := &Config{
		args: []string{".binary", "4"},
		logf: func(f string, a ...interface{}) {
			fmt.Printf(f+"\n", a...)
			logs = append(logs, fmt.Sprintf(f, a...))
		},
		puzzles: testPuzzles(),
	}
	err := Run(config)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "3 out of range, there are only 3 puzzles registered so far")
}

func TestRunFailsWithImpossibleTest(t *testing.T) {
	logs := []string{}
	config := &Config{
		args: []string{".binary", "0"},
		logf: func(f string, a ...interface{}) {
			fmt.Printf(f+"\n", a...)
			logs = append(logs, fmt.Sprintf(f, a...))
		},
		puzzles: testPuzzles(),
	}
	err := Run(config)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "you cannot specify test 0 or under")
}
