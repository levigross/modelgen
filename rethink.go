package main

import (
	"bytes"
	"fmt"
	"go/ast"
)

func generateRethinkDBMethods(file *ast.File) ([]byte, error) {
	exportsFile := &bytes.Buffer{}
	// structsToIgnore := []string{}
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.File:
			astFile := n.(*ast.File)
			exportsFile.WriteString("package " + astFile.Name.String())
			exportsFile.WriteByte('\n')
		case *ast.TypeSpec:
			fmt.Printf("%T %#v\n", x, n)
			mts := n.(*ast.TypeSpec)
			if _, ok := mts.Type.(*ast.StructType); !ok {
				return true
			}
			modelStruct := mts.Type.(*ast.StructType)
			for _, field := range modelStruct.Fields.List {
				fmt.Println(field.Tag.Value)
			}

		case nil:
			return true
		default:
			fmt.Printf("%T %#v\n", x, n)
		}
		return true
	})
	fmt.Println(string(exportsFile.Bytes()))
	return exportsFile.Bytes(), nil
}
