package isUsePointer

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "isUsePointer is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "isUsePointer",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			if IsReceiverPointer(n, info) {
				pass.Reportf(n.Pos(), "use pointer")
			}
		}
	})

	return nil, nil
}
func IsReceiverPointer(method *ast.FuncDecl, info *types.Info) bool {
	if method.Recv == nil {
		return false
	}

	recv := method.Recv.List[0].Type
	recvType := info.TypeOf(recv)
	if recvType == nil {
		return false
	}

	_, isPointer := recvType.Underlying().(*types.Pointer)
	return isPointer
}
