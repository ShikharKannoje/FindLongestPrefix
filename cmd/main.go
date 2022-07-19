package main

import (
	"os"

	"github.com/ShikharKannoje/prefixSearch/pkg"
)

func main() {
	inputStringOfString := os.Args[1:]
	pkg.InitiateSearchFor(inputStringOfString)
}
