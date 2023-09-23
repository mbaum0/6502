package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mbaum0/65c2265/lexer"
)

func main() {

	// gets first CLI are as the file name and reads the file
	// into a string
	fileName := os.Args[1]
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// creates a new lexer and passes the file string to it
	lex := lexer.NewLexer(string(file))
	lex.Lex()

	fmt.Println(lex)
}
