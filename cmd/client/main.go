package main

import (
	"os"

	"rubenkristian.github.com/pinata-client/parser"
)

func main() {
	args := os.Args[1:]

	parser.Parser(args)
}
