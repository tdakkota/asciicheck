package asciicheck

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "asciicheck",
		Doc:      "checks that all code identifiers does not have non-ASCII symbols in the name",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      run,
	}
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		v := n.(*ast.Ident) // n is always an *ast.Ident because we filter for it

		ch, ascii := isASCII(v.Name)
		if !ascii {
			pass.Report(
				analysis.Diagnostic{
					Pos:     v.Pos(),
					Message: fmt.Sprintf("identifier %q contain non-ASCII character: %#U", v.Name, ch),
				},
			)
		}
	})

	return nil, nil
}
