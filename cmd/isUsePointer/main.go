package main

import (
	"isUsePointer"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(isUsePointer.Analyzer) }
