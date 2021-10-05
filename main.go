package main

import (
	"MrTeeny/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
