package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dnnrly/puzzle-template"
)

// Puzzles represents all of the solutions that the application knows about
var Puzzles []puzzle.Puzzle = []puzzle.Puzzle{}

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
	initPuzzles()

	err := Run(&Config{
		args:    os.Args,
		logf:    func(f string, a ...interface{}) { log.Printf(f, a...) },
		puzzles: Puzzles,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// Run executes the command business logic
func Run(c *Config) error {
	c.logf("There are %d solutions\n", len(Puzzles))

	if len(c.args) == 1 {
		for k, s := range c.puzzles {
			start := time.Now()
			n := s()
			total := time.Now().Sub(start).Milliseconds()
			c.logf("%03d %3d - %d\n", total, k+1, n)
		}
	} else if c.args[1] == "latest" {
		start := time.Now()
		i := len(c.puzzles) - 1
		n := c.puzzles[i]()
		total := time.Now().Sub(start).Milliseconds()
		c.logf("%03d %3d - %d\n", total, i+1, n)
	} else {
		for _, k := range os.Args[1:] {
			i, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}

			i--
			if i < 0 || i >= len(c.puzzles) {
				panic(fmt.Sprintf(
					"%d out of range, there are only %d puzzles registered so far\n",
					i,
					len(c.puzzles),
				))
			}

			start := time.Now()
			n := c.puzzles[i]()
			total := time.Now().Sub(start).Milliseconds()
			c.logf("%03d %3d - %d\n", total, i+1, n)
		}
	}

	return nil
}

func initPuzzles() {
	// next puzzle
}
