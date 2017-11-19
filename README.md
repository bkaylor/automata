# automata
Some simple automata simulations in Go

## elementary
`go run elementary.go <integer>` prints the elementary automata with rule <integer> (up to 255).
`go run forest.go` runs a forest fire automata.
`go run life.go` runs a game of life simulation with a random start configuration.
`go run pi.go` lists all integers who use each digit 0-9 exactly once, each substring starting from the right is divisible by its length.

NOTE: pi.go isn't an automata, I don't remember why I wrote it, or why it is called pi.go. Consider it a bonus.
Also- I really wish every other language had a time package that was as easy to use as go's.
