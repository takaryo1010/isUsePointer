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

type isUsePointerResult struct {
	count_use     bool
	count_not_use bool
	filename      string
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	info := pass.TypesInfo
	var res isUsePointerResult
	new_fname := pass.Fset.File(pass.Files[0].Pos()).Name()
	if res.filename != new_fname {
		res.filename = new_fname
		res.count_not_use = false
		res.count_use = false
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			switch IsReceiverPointer(n, info) {

			case 1:
				res.count_use = true
				if res.count_not_use {
					pass.Reportf(n.Pos(), "use pointer & Mixed use and non-use of pointers")
				} else {
					pass.Reportf(n.Pos(), "use pointer")
				}

			case 2:
				res.count_not_use = true
				if res.count_use {
					pass.Reportf(n.Pos(), "not use pointer & Mixed use and non-use of pointers")
				} else {
					pass.Reportf(n.Pos(), "not use pointer")
				}

			}
		}
	})
	return nil, nil
}
func IsReceiverPointer(method *ast.FuncDecl, info *types.Info) int {
	if method.Recv == nil {
		return 0
	}

	recv := method.Recv.List[0].Type
	recvType := info.TypeOf(recv)
	if recvType == nil {
		return 0
	}

	_, isPointer := recvType.Underlying().(*types.Pointer)
	if isPointer {
		return 1
	}
	return 2
}
