# Puzzle Template

This is a template that will allow you to solve programming puzzles with out worrying about the mechanics of managing a project.

For more information about the kind of delightful puzzles I'm talking about, you should visit https://projecteuler.net and
https://adventofcode.com/. It's fun, really, trust me!

This assumes that you have the following installed:
* [Go](https://go.dev/)
* [Make](https://en.wikipedia.org/wiki/Make_(software))

You can run this template anywhere you can install these tools (technically, you don't even need make). I've even run this on phone with [Termux](https://termux.com/)! There's no excuse for not giving it a go. :)

If you want to know more about Go then a great place to start is the [Go Tour](https://tour.golang.org) followed quickly by [Effective Go](https://golang.org/doc/effective_go.html) and [Go By Example](https://gobyexample.com/).

There are some helpers in [puzzle.go] - but I've not made it TOO easy, you will have to fill these out yourself. Maybe you could extend them in interesting ways? What if you had a nice way of caching prime numbers perhaps? I'll leave it up to you.

This project is based my original template for Project Euler that can be found at https://github.com/dnnrly/euler-template.

## Creating your own version

To solve the puzzles you will need to fork this project in to your own user area. From there, clone it to a local computer and you're ready to go. You can find istructions on how to do this in Github itself.

If there are ANY problems with this template then feel free to raise an issue and I'll see what I can do. But beware, "I can't figure out the solution with puzzle X" does not constitute a problem with this template. Chances are that I haven't solved that one either. To be honest, I've only solved the first 10 anyway.

## Running your puzzles

All of these steps assume that you have Make installed. If you don't you can check out the contents of `Makefile` to see the real shell command being run for each target.

To run all of your puzzles:
```bash
$ make run-all
```

To run only the last test:
```bash
$ make latest
```

To run a single other test:
```bash
$ go run cmd/puzzle/main.go 4
```
Or you could use a different number. 4 is my favourite. :)

## Adding the next puzzle

It is possible to generate a new file for your next puzzle using the following command from the root of the project:
```bash
$ make next
```

After this, it's up to you to solve the puzzle. The thing that's just been generated will look a little like this:
```go
func Puzzle001() Puzzle {
    // In this function, you can add large strings (your input data, perhaps) so
    // that they can be parsed in the initialisation phase of the puzzle and won't
    // count to the time taken for your algorithm.
    data := `1
2
3`

    // You can add multiple variables that you use later on if you wish.
    values := []int{}

    // And you can even add helpers that can be called only from inside your
    // puzzle solution.
    convert := func(s string) int {
        v, _ := strconv.Atoi(s)
        return v
    }

    return Puzzle{
        // Init will be called before any of the solutions. Do all of your
        // expsensive pre-processing here.
        Init: func() {
            lines := strings.Split(data, "\n")
            for _, l := range lines {
                values = append(values, convert(l))
            }
        },
		// Parts contains all of the different sub-solutions that you need
		// to implement (looking at you Advent of Code).
        Parts: []Solution{
            // You can consume the pre-processed data in the solutions here
            func() int { return values[0] + 100 },
            func() int { return values[1] + 100 },
            func() int { return values[2] + 100 },
        },
    }
}
```

Once you've got the answers you need, I would recommend something like the following:
```bash
$ git add .
$ git commit -m "Add solution to puzzle $(ls -l | grep puzzle | wc -l)"
$ git push
```

## Adding tests for your puzzles

Sometimes it's necessary to check that your code is doing what you expect without actually solving a puzzle. For example, you might want to check that a new prime number checker algorithm works. In this case, you would simply add a `_test.go` file the you would do normally and add your test code there. To execute the tests, use the following:
```bash
$ make test
```

Just beware, if you create your helper functions inside your generated puzzle function then they won't be accessible from your tests. To get around this, just create them outside of this function. And remember that `common.go` has been created for those helpers that you think might be used by several puzzles.

## Contributing to this template

Feel free to fork this repository and raise a pull request for any changes.

To run the on the framework as well in the puzzle package, you will need to do the following:
```bash
$ make test test-all acceptance-test
```

The last target will run the `next` and `puzzle` commands, making sure that the generated code builds and runs - it doesn't do any fancy validation.

This will run the tests for the different individual commands.
