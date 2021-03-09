package main

import (
	"github.com/chkhetiani/markov/input"
	"github.com/chkhetiani/markov/output"
	"github.com/chkhetiani/markov/solver"
)

func main() {
	data := input.Parse()
	result := solver.Solve(data)
	output.Print(result)
}
