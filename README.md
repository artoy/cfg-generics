# cfg-generics
`cfg-generics` is CFG (Control Flow Graph) generator from some information in your compiler written by Go.

All you need is...

- head instructions of function
- a function returns predecessors of each instruction
- a function returns successors of each instruction

## Note
`cfg-generics` is **pre-release**. In this version, you can only generate and dump CFG. More functions such as analysis or optimize with CFG, are going to be added in next version.

## Requirement
Go 1.18+

## Installation

`go install https://github.com/artoy/cfg-generics@latest`


## Usage
Below is an example generate CFG with `cfg-generics`. Use `Print` function to dump generated CFG.

```go
package main

import (
	"github.com/artoy/cfg-generics"
	"fmt"
)

type Inst struct {
	Preds   []*Inst
	Content string
	Succs   []*Inst
}

func main() {
	aDecl := &Inst{
		Preds:   nil,
		Content: "var a = 10",
		Succs:   nil,
	}

	plus1Inst := &Inst{
		Preds:   []*Inst{aDecl},
		Content: "a++",
		Succs:   nil,
	}

	ifStmt := &Inst{
		Preds:   []*Inst{plus1Inst},
		Content: "if a > 5",
		Succs:   nil,
	}

	thenClause := &Inst{
		Preds:   []*Inst{ifStmt},
		Content: "fmt.Println(\"a is big!\")",
		Succs:   nil,
	}

	elseClause := &Inst{
		Preds:   []*Inst{ifStmt},
		Content: "fmt.Println(\"a is small...\")",
		Succs:   nil,
	}

	printStmt := &Inst{
		Preds:   []*Inst{thenClause, elseClause},
		Content: "fmt.Println(\"end\")",
		Succs:   nil,
	}

	returnStmt := &Inst{
		Preds:   []*Inst{printStmt},
		Content: "return",
		Succs:   nil,
	}

	aDecl.Succs = []*Inst{plus1Inst}
	plus1Inst.Succs = []*Inst{ifStmt}
	ifStmt.Succs = []*Inst{thenClause, elseClause}
	thenClause.Succs = []*Inst{printStmt}
	elseClause.Succs = []*Inst{printStmt}
	printStmt.Succs = []*Inst{returnStmt}

	hs := []*Inst{aDecl}
	preds := func(i *Inst) []*Inst {
		return i.Preds
	}
	succs := func(i *Inst) []*Inst {
		return i.Succs
	}

	f := cfg.NewFunction[Inst](hs, preds, succs)

	p := func(i *Inst) {
		fmt.Println(i.Content)
	}
	for _, v := range f.CFGs {
		v.Print(p)
	}
}
```
