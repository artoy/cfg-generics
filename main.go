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
	aDecl := Inst{
		Preds:   nil,
		Content: "var a = 10",
		Succs:   nil,
	}

	ifStmt := Inst{
		Preds:   []*Inst{&aDecl},
		Content: "if a > 5",
		Succs:   nil,
	}

	thenClause := Inst{
		Preds:   []*Inst{&ifStmt},
		Content: "fmt.Println(\"Big!\")",
		Succs:   nil,
	}

	elseClause := Inst{
		Preds:   []*Inst{&ifStmt},
		Content: "fmt.Println(\"Small...\")",
		Succs:   nil,
	}

	returnStmt := Inst{
		Preds:   []*Inst{&thenClause, &elseClause},
		Content: "fmt.Println(\"Small...\")",
		Succs:   nil,
	}

	aDecl.Succs = []*Inst{&ifStmt}
	ifStmt.Succs = []*Inst{&thenClause, &elseClause}
	thenClause.Succs = []*Inst{&returnStmt}
	elseClause.Succs = []*Inst{&returnStmt}

	hs := []*Inst{&aDecl}
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
