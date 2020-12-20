package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	src := "cmd/euler/main.go"

	f, err := os.Open(src)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	defer f.Close()

	contents, err := ioutil.ReadAll(f)
	re := regexp.MustCompile("Puzzle\\d+")
	all := re.FindAll(contents, -1)
	next := len(all) + 1

	text := fmt.Sprintf("\tSolutions = append(Solutions, euler.Puzzle%03d)\n\t// next puzzle", next)

	re = regexp.MustCompile("\t// next puzzle")
	update := re.ReplaceAllString(string(contents), text)
	err = ioutil.WriteFile(src, []byte(update), 0660)
	if err != nil {
		log.Fatalf("unable to write main file: %v", err)
	}

	puzzle := fmt.Sprintf(`package euler
	
func Puzzle%03d() int {
	return 0
}`, next)

	err = ioutil.WriteFile(
		fmt.Sprintf("puzzle%03d.go", next),
		[]byte(puzzle),
		0660,
	)
	if err != nil {
		log.Fatalf("unable to write puzzle file: %v", err)
	}

	fmt.Printf("Created Puzzle%03d\n", next)
}
