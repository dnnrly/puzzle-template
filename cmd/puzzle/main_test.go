package main

import (
	"fmt"
	"testing"

	"github.com/dnnrly/puzzle-template"
	"github.com/stretchr/testify/assert"
)

func TestRunsWithtoutError(t *testing.T) {
	logs := []string{}
	config := &Config{
		args: []string{".binary"},
		logf: func(f string, a ...interface{}) {
			fmt.Printf(f, a...)
			logs = append(logs, fmt.Sprintf(f, a...))
		},
		puzzles: []puzzle.Puzzle{},
	}
	err := Run(config)

	assert.NoError(t, err)
}
