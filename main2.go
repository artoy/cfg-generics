package cfgGenerics

//
//import (
//	"flag"
//	"fmt"
//	"golang.org/x/tools/go/packages"
//	"golang.org/x/tools/go/ssa"
//	"golang.org/x/tools/go/ssa/ssautil"
//	"os"
//)
//
//func main() {
//	flag.Parse()
//	if err := run(); err != nil {
//		fmt.Fprintln(os.Stderr, err)
//		os.Exit(1)
//	}
//}
//
//func run() error {
//	//if len(os.Args) < 2 {
//	//	return errors.New("パターンを指定してください")
//	//}
//
//	cfg := &packages.Config{
//		Mode: packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
//	}
//	pkgs, err := packages.Load(cfg, "./test.go")
//	if err != nil {
//		return err
//	}
//
//	_, ssaPkgs := ssautil.Packages(pkgs, ssa.BuilderMode(0))
//
//	for _, pkg := range ssaPkgs {
//		if err := analyze(pkg); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func analyze(pkg *ssa.Package) error {
//	for _, v := range pkg.Members {
//		switch v.(type) {
//		case *ssa.Function:
//			//v.(*ssa.Function).Blocks
//		}
//	}
//
//	//inspect := inspector.New(pkg.Syntax)
//	//
//	//nodeFilter := []ast.Node{
//	//	(*ast.FuncDecl)(nil),
//	//}
//	//
//	//inspect.Preorder(nodeFilter, func(n ast.Node) {
//	//	switch n := n.(type) {
//	//	case *ast.FuncDecl:
//	//		it := make([]*_go.GoIR, 0)
//	//		_go.Translation(n.Body, it)
//	//	}
//	//})
//
//	return nil
//}
