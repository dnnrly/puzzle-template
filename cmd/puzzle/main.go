package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dnnrly/puzzle-template"
)

// Logger sends log entries to a real logger
type Logger func(f string, a ...interface{})

// Config contains all of the different things that might change between tests and real main
type Config struct {
	args    []string
	logf    Logger
	puzzles []puzzle.Puzzle
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	err := Run(&Config{
		args:    os.Args,
		logf:    func(f string, a ...interface{}) { log.Printf(f, a...) },
		puzzles: puzzle.RegisterPuzzles(),
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// Run executes the command business logic
func Run(c *Config) error {
	c.logf("There are %d solutions", len(c.puzzles))

	if len(c.args) == 1 {
		for k, s := range c.puzzles {
			runPuzzle(k+1, s, c.logf)
		}
	} else if c.args[1] == "latest" {
		i := len(c.puzzles) - 1
		runPuzzle(i+1, c.puzzles[i], c.logf)
	} else {
		for _, k := range c.args[1:] {
			i, err := strconv.Atoi(k)
			if err != nil {
				return fmt.Errorf("'%s' is not an option", k)
			}

			if i <= 0 {
				return fmt.Errorf("you cannot specify test 0 or under")
			}

			i--
			if i >= len(c.puzzles) {
				return fmt.Errorf(
					"%d out of range, there are only %d puzzles registered so far",
					i,
					len(c.puzzles),
				)
			}

			runPuzzle(i+1, c.puzzles[i], c.logf)
		}
	}

	return nil
}

func runPuzzle(k int, puz puzzle.Puzzle, logf Logger) {
	for i, p := range puz.Parts {
		start := time.Now()
		n := p()
		total := time.Now().Sub(start).Milliseconds()
		logf("%03d %3d.%d - %d", total, k, i+1, n)
	}
}
