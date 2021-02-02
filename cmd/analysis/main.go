package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func main() {
	group := os.Args[1]
	kind := os.Args[2]
	suffix := "WithContext"
	fset := token.NewFileSet()
	node, err := parser.ParseDir(fset, fmt.Sprintf("/Users/monus/go/src/github.com/crossplane/provider-aws/pkg/controller/%s/%s", group, kind), nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	calls := map[string]bool{}

	ast.Inspect(node[kind], func(x ast.Node) bool {
		s, ok := x.(*ast.CallExpr)
		if !ok {
			return true
		}
		name := ""
		switch expr := s.Fun.(type) {
		case *ast.SelectorExpr:
			name = expr.Sel.Name
		case *ast.Ident:
			name = expr.Name
		default:
			return false
		}
		if strings.HasSuffix(name, suffix) {
			actionName := fmt.Sprintf("%s:%s", group, strings.Split(name, suffix)[0])
			calls[actionName] = true
		}
		return true
	})
	result := map[string][]string{}
	for c := range calls {
		result["Action"] = append(result["Action"], c)
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
