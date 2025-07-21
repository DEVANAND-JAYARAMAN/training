package main

import (
	"fmt"
)

// Define the interface
type Match interface {
	Innings()
}

// Structs
type First struct {
	wicket, run int
}

type Second struct {
	wicket, run int
}

// Implementing the interface for First
func (f First) Innings() {
	fmt.Println("Run: 100")
	fmt.Println("Wickets: 04")
}

// Implementing the interface for Second
func (s Second) Innings() {
	fmt.Println("Runs: 200")
	fmt.Println("Wickets: 06")
}

func main() {
	var h Match

	h = First{}
	h.Innings()

	h = Second{}
	h.Innings()
}
