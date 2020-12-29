package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/spf13/afero"
)

// Config contains all of the different things that might change between tests and real main
type Config struct {
	fs afero.Fs
}

func main() {
	err := Run(&Config{
		fs: afero.NewOsFs(),
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// Run executes the command business logic
func Run(config *Config) error {
	src := "puzzles.go"

	f, err := config.fs.Open(src)
	if err != nil {
		return fmt.Errorf("unable to read main file: %w", err)
	}
	defer f.Close()

	contents, err := ioutil.ReadAll(f)
	re := regexp.MustCompile("Puzzle\\d+")
	all := re.FindAll(contents, -1)
	next := len(all) + 1

	text := fmt.Sprintf("\tpuzzles = append(puzzles, Puzzle%03d())\n\t// next puzzle", next)

	re = regexp.MustCompile("\t// next puzzle")
	update := re.ReplaceAllString(string(contents), text)
	err = afero.WriteFile(config.fs, src, []byte(update), 0660)
	if err != nil {
		return fmt.Errorf("unable to write main file: %w", err)
	}

	puzzle := fmt.Sprintf(`package puzzle

// Puzzle%03d is the solutions for puzzle %d
func Puzzle%03d() Puzzle {
	// Here's where you put some extra data you might need to process
	data := ""

	// You can create helpers like this that can be called later
	processData := func() {
		data += "a"
	}

	return Puzzle {
		// Init will be called before any of the solutions. Do all of your
		// expsensive pre-processing here.
		Init: func() {
			processData()
		},
		// Parts contains all of the different sub-solutions that you need
		// to implement (looking at you Advent of Code).
		Parts: []Solution{
			func() int { return 0 },
		},
	}
}`, next, next, next)

	err = afero.WriteFile(
		config.fs,
		fmt.Sprintf("puzzle%03d.go", next),
		[]byte(puzzle),
		0660,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Created Puzzle%03d\n", next)

	return nil
}
