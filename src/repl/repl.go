package repl

import (
	"bufio"
	"fmt"
	"io"

	"ayush.interpreter.monkey/src/lexer"
	"ayush.interpreter.monkey/src/token"
)

const PROMT = ">>"

/* REPL - Read Eval Print Loop
 */
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for { // start infinite loop to get input from console.
		fmt.Fprintf(out, PROMT)   // >> ...
		scanned := scanner.Scan() // waits for user input
		if !scanned {
			return
		}

		line := scanner.Text() // Gets the input line

		l := lexer.New(line) // initialize lexer with our input line...

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok) // repeatedly call NextToken until EOF is hitf
		}
	}
}
