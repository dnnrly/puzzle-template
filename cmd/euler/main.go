package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dnnrly/euler"
)

// Solutions represents all of the solutions that the application knows about
var Solutions []euler.Solution = []euler.Solution{}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	initPuzzles()

	fmt.Printf("There are %d solutions\n", len(Solutions))

	if len(os.Args) == 1 {
		for k, s := range Solutions {
			start := time.Now()
			n := s()
			total := time.Now().Sub(start).Milliseconds()
			log.Printf("%03d %3d - %d\n", total, k+1, n)
		}
	} else if os.Args[1] == "latest" {
		start := time.Now()
		i := len(Solutions) - 1
		n := Solutions[i]()
		total := time.Now().Sub(start).Milliseconds()
		log.Printf("%03d %3d - %d\n", total, i+1, n)
	} else {
		for _, k := range os.Args[1:] {
			i, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}

			i--
			if i < 0 || i >= len(Solutions) {
				panic(fmt.Sprintf(
					"%d out of range, there are only %d puzzles registered so far\n",
					i,
					len(Solutions),
				))
			}

			start := time.Now()
			n := Solutions[i]()
			total := time.Now().Sub(start).Milliseconds()
			log.Printf("%03d %3d - %d\n", total, i+1, n)
		}
	}
}

func initPuzzles() {
	// next puzzle
}
