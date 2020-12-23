package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dnnrly/puzzle-template"
)

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
		puzzles: []puzzle.Puzzle{
			func() int { return 1 },
			func() int { return 2 },
			func() int { return 3 },
		},
	}
	err := Run(config)

	assert.NoError(t, err)
	assert.Contains(t, logs, "There are 3 solutions")
	assert.Contains(t, logs, "000   1 - 1")
	assert.Contains(t, logs, "000   2 - 2")
	assert.Contains(t, logs, "000   3 - 3")
}

func TestRunsLatestWithtoutError(t *testing.T) {
	logs := []string{}
	config := &Config{
		args: []string{".binary", "latest"},
		logf: func(f string, a ...interface{}) {
			t.Logf(f+"\n", a...)
			logs = append(logs, fmt.Sprintf(f, a...))
		},
		puzzles: []puzzle.Puzzle{
			func() int { return 1 },
			func() int { return 2 },
			func() int { return 3 },
		},
	}
	err := Run(config)

	assert.NoError(t, err)
	assert.Contains(t, logs, "There are 3 solutions")
	assert.NotContains(t, logs, "000   1 - 1")
	assert.NotContains(t, logs, "000   2 - 2")
	assert.Contains(t, logs, "000   3 - 3")
}

func TestRunsSpecificTest(t *testing.T) {
	logs := []string{}
	config := &Config{
		args: []string{".binary", "2"},
		logf: func(f string, a ...interface{}) {
			t.Logf(f+"\n", a...)
			logs = append(logs, fmt.Sprintf(f, a...))
		},
		puzzles: []puzzle.Puzzle{
			func() int { return 1 },
			func() int { return 2 },
			func() int { return 3 },
		},
	}
	err := Run(config)

	assert.NoError(t, err)
	assert.Contains(t, logs, "There are 3 solutions")
	assert.NotContains(t, logs, "000   1 - 1")
	assert.Contains(t, logs, "000   2 - 2")
	assert.NotContains(t, logs, "000   3 - 3")
}

func TestRunFailsWithUnparsableParameters(t *testing.T) {
	logs := []string{}
	config := &Config{
		args: []string{".binary", "rubbish"},
		logf: func(f string, a ...interface{}) {
			fmt.Printf(f+"\n", a...)
			logs = append(logs, fmt.Sprintf(f, a...))
		},
		puzzles: []puzzle.Puzzle{
			func() int { return 1 },
			func() int { return 2 },
			func() int { return 3 },
		},
	}
	err := Run(config)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "'rubbish' is not an option")
}
