package asciicheck

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "asciicheck",
		Doc:  "checks that all code identifiers does not have non-ASCII symbols in the name",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
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
	if v, ok := n.(*ast.Ident); ok {
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
}

func getIdentifiers(file *ast.File) (idents []*ast.Ident) {
	for _, decl := range file.Decls {
		switch value := decl.(type) {
		case *ast.GenDecl:
			idents = append(idents, handleGenDecl(value)...)
		case *ast.FuncDecl:
			idents = append(idents, handleFuncDecl(value)...)
		}
	}

	return idents
}

func handleGenDecl(d *ast.GenDecl) []*ast.Ident {
	fmt.Println("gendecl:", d)

	var r []*ast.Ident
	for _, spec := range d.Specs {
		switch value := spec.(type) {
		case *ast.ValueSpec:
			fmt.Println("\tspec_value:", value, len(value.Values))
			for _, name := range value.Names {
				r = append(r, name)
				fmt.Println("\t\tname:", name)
			}

			for _, expr := range value.Values {
				fmt.Println(expr)
			}
		case *ast.TypeSpec:
			r = append(r, value.Name)
			fmt.Println("\tspec_type:", value.Name)
		}
	}

	return r
}

func handleFuncDecl(d *ast.FuncDecl) []*ast.Ident {
	fmt.Println("funcdecl:", d)

	r := []*ast.Ident{d.Name}
	for _, stmt := range d.Body.List {
		switch v := stmt.(type) {
		case *ast.AssignStmt:
			for _, value := range v.Lhs {
				if v, ok := value.(*ast.Ident); ok {
					r = append(r, v)
				}
			}

			for _, value := range v.Rhs {
				if v, ok := value.(*ast.Ident); ok {
					r = append(r, v)
				}
			}
		}
	}

	return r
}
