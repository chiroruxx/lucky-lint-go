package main

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"
)

var analyzer = &analysis.Analyzer{
	Name:     "luckyLint",
	Doc:      "lint your naming is lucky",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var thresholdLevel = luckyLevelBest

var reservedFuncNames = map[string]bool{
	"init": true,
	"main": true,
}

func main() {
	singlechecker.Main(analyzer)
}

func run(pass *analysis.Pass) (any, error) {
	var filter []ast.Node
	skips := make(map[ast.Node]bool)

	i := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	i.Preorder(filter, func(n ast.Node) {
		if skips[n] {
			return
		}

		switch n := n.(type) {
		// var, const
		case *ast.ValueSpec:
			for _, name := range n.Names {
				if name.Name == "ok" {
					return
				}
				assert(name.Name, name.Pos(), pass)
			}
		// const
		case *ast.TypeSpec:
			assert(n.Name.Name, n.Name.Pos(), pass)
		// :=
		case *ast.AssignStmt:
			for _, name := range n.Lhs {
				ident, ok := name.(*ast.Ident)
				if !ok {
					continue
				}
				if ident.Name == "ok" {
					return
				}
				assert(ident.Name, ident.Pos(), pass)
			}
		// func
		case *ast.FuncDecl:
			if n.Recv == nil {
				// skip reserved func
				if reservedFuncNames[n.Name.Name] {
					return
				}
			} else {
				// skip receiver
				for _, field := range n.Recv.List {
					skips[field] = true
				}
			}
			assert(n.Name.Name, n.Name.Pos(), pass)
		// struct field
		case *ast.Field:
			for _, name := range n.Names {
				assert(name.Name, name.Pos(), pass)
			}
		// import
		case *ast.ImportSpec:
			if n.Name == nil {
				return
			}
			if n.Name.Name == "." {
				return
			}
			assert(n.Name.Name, n.Name.Pos(), pass)
		}
	})
	return nil, nil
}

func assert(name string, pos token.Pos, pass *analysis.Pass) {
	if name == "_" {
		return
	}

	lv := divine(name)
	if lv >= thresholdLevel {
		return
	}

	pass.Reportf(pos, "naming '%s' is not lucky", name)
}
