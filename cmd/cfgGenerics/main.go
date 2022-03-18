package main

import (
	cfgGenerics "cfg-generics"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(cfgGenerics.Analyzer) }
