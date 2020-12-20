# Euler Template

This is a template that will allow you to solve Project Euler puzzles with out worrying about the mechanics of managing a project.

For more information about these delightful puzzles, you need to visit https://projecteuler.net - it's fun, really, trust me!

This assumes that you have the following installed:
* [Go](https://go.dev/)
* [Make](https://en.wikipedia.org/wiki/Make_(software))

You can run this template anywhere you can install these tools (technically, you don't even need make). I've even run this on phone with [Termux](https://termux.com/)! There's no excuse for not giving it a go. :)

If you want to know more about Go then a great place to start is the [Go Tour](https://tour.golang.org) followed quickly by [Effective Go](https://golang.org/doc/effective_go.html) and [Go By Example](https://gobyexample.com/).

There are some helpers in [euler.go] - but I've not made it TOO easy, you will have to fill these out yourself. Maybe you could extend them in interesting ways? What if you had a nice way of caching prime numbers perhaps? I'll leave it up to you.

## Creating your own version

To solve the puzzles you will need to fork this project in to your own user area. From there, clone it to a local computer and you're ready to go. You can find istructions on how to do this in Github itself.

If there are ANY problems with this template then feel free to raise an issue and I'll see what I can do. But beware, "I can't figure out the solution with puzzle X" does not constitute a problem with this template. Chances are that I haven't solved that one either. To be honest, I've only solved the first 10 anyway.

## Running your puzzles

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
$ go run cmd/make/main.go 4
```
Or you could use a different number. 4 is my favourite. :)

## Adding the next puzzle

It is possible to generate a new file for your next puzzle using the following command from the root of the project:
```bash
$ go run cmd/next/main.go
```

After this, it's up to you to solve the puzzle. I would recommend something like the following when you figure out the answer:
```bash
$ git add .
$ git commit -m "Add solution to puzzle $(ls -l | grep puzzle | wc -l)"
$ git push
```