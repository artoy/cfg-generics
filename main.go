package main

import (
	"cfg-generics/cfg"
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
		Content: "fmt.Println(\"return 0\")",
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
