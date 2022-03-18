package _go

import "go/ast"

type GoIR struct {
	Inst  ast.Stmt
	Preds []*GoIR
	Succs []*GoIR
}

func (i GoIR) Print() {
}
