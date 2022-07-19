package main

import (
	"os"

	"github.com/ShikharKannoje/prefixSearch/pkg"
)

func main() {
	// reading the commandline arguments
	inputStringOfString := os.Args[1:]
	pkg.InitiateSearchFor(inputStringOfString)
}
