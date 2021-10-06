package main

import (
	"MrTeeny/repl"
	"os"
)

func main() {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;
if (5 < 10) {
return true;
} else {
return false;
}`
	repl.Start(os.Stdin, os.Stdout, input)
}
