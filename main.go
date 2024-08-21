package main

import "fmt"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	fmt.Println("Current input:", l.input)
	l.readChar()
	return l
}

// readChar advances the lexer to the next character in the input string.
// It updates the current character (ch) and positions (position and readPosition).
// If the lexer has reached the end of the input, readChar sets ch to 0 ('NUL'),
// indicating that there are no more characters to process.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for 'NUL', indicating end of input
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func main() {
	// Step 1: Initialize a new Lexer with some input string
	input := "let x = 10;"
	lexer := New(input)

	// Step 2: Print the initial state of the Lexer
	fmt.Println("Initial Lexer State:")
	fmt.Printf("Input: %s\n", lexer.input)
	// fmt.Printf("Position: %d\n", lexer.position)
	// fmt.Printf("ReadPosition: %d\n", lexer.readPosition)
	// fmt.Printf("Current char: %c\n", lexer.ch)

	// Step 3: Simulate reading characters one by one
	// fmt.Println("\nReading characters...")
	// for lexer.ch != 0 { // Continue until we reach the end of the input
	// 	fmt.Printf("Current char: %c, Position: %d, ReadPosition: %d\n", lexer.ch, lexer.position, lexer.readPosition)
	// 	lexer.readChar()
	// }

	// Step 4: Final state of the Lexer after processing the input
	// fmt.Println("\nFinal Lexer State:")
	// fmt.Printf("Position: %d\n", lexer.position)
	// fmt.Printf("ReadPosition: %d\n", lexer.readPosition)
	// fmt.Printf("Current char (should be NUL): %c\n", lexer.ch)
}
