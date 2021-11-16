package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cian911/platoon/lexer"
	"github.com/cian911/platoon/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		// Get the new latest line and pass it to our lexer
		line := scanner.Text()

		if line == "exit" {
			fmt.Print("Cheers, thanks for coming.\n")
			return
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
