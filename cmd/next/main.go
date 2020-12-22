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
	src := "cmd/puzzle/main.go"

	f, err := config.fs.Open(src)
	if err != nil {
		return fmt.Errorf("unable to read main file: %w", err)
	}
	defer f.Close()

	contents, err := ioutil.ReadAll(f)
	re := regexp.MustCompile("Puzzle\\d+")
	all := re.FindAll(contents, -1)
	next := len(all) + 1

	text := fmt.Sprintf("\tSolutions = append(Solutions, puzzle.Puzzle%03d)\n\t// next puzzle", next)

	re = regexp.MustCompile("\t// next puzzle")
	update := re.ReplaceAllString(string(contents), text)
	err = afero.WriteFile(config.fs, src, []byte(update), 0660)
	if err != nil {
		return fmt.Errorf("unable to write main file: %w", err)
	}

	puzzle := fmt.Sprintf(`package puzzle
	
func Puzzle%03d() int {
	return 0
}`, next)

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
