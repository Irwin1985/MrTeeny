package repl

import (
	"MrTeeny/lexer"
	"MrTeeny/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		tok := l.NextToken()

		for tok.Type != token.EOF {
			fmt.Print(tok.ToString() + "\n")
			tok = l.NextToken()
		}
		fmt.Print(tok.ToString() + "\n")
	}
}
