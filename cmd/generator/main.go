package main

import (
	"fmt"
	typewriter "github.com/muvaf/typewriter/pkg/traverser"
	"github.com/muvaf/typewriter/pkg/wrapper"
	"go/types"
	"golang.org/x/tools/go/packages"
)

const (
	LoadMode = packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax
)

func main() {
	localDynamoDBPkg := "github.com/crossplane/provider-aws/apis/dynamodb/v1alpha1"
	remoteDynamoDBPkg := "github.com/aws/aws-sdk-go/service/dynamodb"
	pkgs, err := packages.Load(&packages.Config{Mode: LoadMode}, localDynamoDBPkg, remoteDynamoDBPkg)
	if err != nil {
		panic(err)
	}
	fmt.Println("package loading completed")
	var aPkg *types.Package
	var bPkg *types.Package
	for _, p := range pkgs {
		if p.Name == "v1alpha1" {
			aPkg = p.Types
		}
		if p.Name == "dynamodb" {
			bPkg = p.Types
		}
	}
	if err := Run(aPkg, bPkg); err != nil {
		panic(err)
	}
}

func Run(p1, p2 *types.Package) error {
	fl := wrapper.NewFile(
		wrapper.WithHeaderPath("/Users/monus/go/src/github.com/crossplane/provider-aws/hack/boilerplate.go.txt"),
	)
	tr := typewriter.NewType(fl.Imports)
	a := p1.Scope().Lookup("TableParameters").Type().(*types.Named)
	b := p2.Scope().Lookup("CreateTableInput").Type().(*types.Named)

	funcContent, err := tr.Print(a, b, "a", "b", 0)
	if err != nil {
		return err
	}
	fn := wrapper.NewFunc(fl.Imports)
	obj, err := fn.Wrap(a, b, "GenerateCreateTableInputNew", funcContent)
	if err != nil {
		return err
	}

	file, err := fl.Wrap("newpkg", obj)
	if err != nil {
		return err
	}
	fmt.Print(file)
	return nil
}
