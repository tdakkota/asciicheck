package asciicheck

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "asciicheck",
		Doc:  "checks that all code identifiers does not have non-ASCII symbols in the name",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(
			file, func(node ast.Node) bool {
				cb(pass, node)
				return true
			},
		)
	}
	return nil, nil
}

func cb(pass *analysis.Pass, n ast.Node) {
	switch n := n.(type) {
	case *ast.File:
		checkIdent(pass, n.Name)
	case *ast.ImportSpec:
		checkIdent(pass, n.Name)
	case *ast.TypeSpec:
		checkIdent(pass, n.Name)
		checkFieldList(pass, n.TypeParams)
	case *ast.ValueSpec:
		for _, name := range n.Names {
			checkIdent(pass, name)
		}
	case *ast.FuncDecl:
		checkIdent(pass, n.Name)
		checkFieldList(pass, n.Recv)
	case *ast.StructType:
		checkFieldList(pass, n.Fields)
	case *ast.FuncType:
		checkFieldList(pass, n.TypeParams)
		checkFieldList(pass, n.Params)
		checkFieldList(pass, n.Results)
	case *ast.InterfaceType:
		checkFieldList(pass, n.Methods)
	case *ast.LabeledStmt:
		checkIdent(pass, n.Label)
	case *ast.AssignStmt:
		if n.Tok == token.DEFINE {
			for _, expr := range n.Lhs {
				if ident, ok := expr.(*ast.Ident); ok {
					checkIdent(pass, ident)
				}
			}
		}
	}
}

func checkIdent(pass *analysis.Pass, v *ast.Ident) {
	if v == nil {
		return
	}

	ch, ascii := isASCII(v.Name)
	if !ascii {
		pass.Report(
			analysis.Diagnostic{
				Pos:     v.Pos(),
				Message: fmt.Sprintf("identifier \"%s\" contain non-ASCII character: %#U", v.Name, ch),
			},
		)
	}
}

func checkFieldList(pass *analysis.Pass, f *ast.FieldList) {
	if f == nil {
		return
	}
	for _, f := range f.List {
		for _, name := range f.Names {
			checkIdent(pass, name)
		}
	}
}
