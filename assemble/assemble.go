package assemble

import (
	"bytes"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const pjName = "go_kyopro"

var (
	gopath     string
	libPath    string
	kyoproName = "kyopro"
	kyoproPath string

	shell string
)

func init() {
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("env GOPATH is not set")
	}
	libPath = filepath.Join(gopath, "src")
	shell = os.Getenv("SHELL")
}

func Assemble(cmd *cobra.Command, filePath string) error {
	cb, err := cmd.PersistentFlags().GetBool("clipBoard")
	if err != nil {
		return err
	}
	fset := token.NewFileSet()
	main, err := parser.ParseFile(fset, filePath, nil, parser.Mode(0))
	if err != nil {
		return err
	}

	IsUsedLib := false
	var imports *[]ast.Spec
	for _, d := range main.Decls {
		if i, ok := d.(*ast.GenDecl); ok {
			imports = &i.Specs
		}
	}
	for _, i := range main.Imports {
		if strings.Contains(i.Path.Value, pjName) {
			if i.Name != nil {
				kyoproName = i.Name.Name
			}
			kyoproPath = i.Path.Value
			kyoproPath = strings.Replace(kyoproPath, "\"", "", -1)
			IsUsedLib = true
		}
	}
	if !IsUsedLib { //ライブラリの参照なし
		return output(main, cb)
	}
	libPath = filepath.Join(libPath, kyoproPath)
	libFiles, err := ioutil.ReadDir(libPath)
	for _, file := range libFiles {
		if file.IsDir() || strings.HasSuffix(file.Name(), "_test.go") {
			continue
		}

		fileName := filepath.Join(libPath, file.Name())
		err = libAppend(&main.Decls, imports, fileName)
		if err != nil {
			return err
		}
	}

	return output(main, cb)
}

func output(out *ast.File, cb bool) error {
	var buf bytes.Buffer
	err := format.Node(&buf, token.NewFileSet(), out)
	if err != nil {
		return err
	}

	outStr := buf.String()
	outStr = strings.Replace(outStr, kyoproName+".", "", -1)
	outStr = strings.Replace(outStr, `"`, `\"`, -1)

	cmdStr := fmt.Sprintf("echo \"%s\"| goimports", outStr)
	formatted, err := exec.Command(shell, "-c", cmdStr).Output()
	if err != nil {
		return err
	}
	outStr = string(formatted)
	if cb {
		err = clipboard.WriteAll(outStr)
		if err != nil {
			return err
		}
	} else {
		fmt.Println(outStr)
	}
	return nil
}

func libAppend(decls *[]ast.Decl, imports *[]ast.Spec, fileName string) error {
	fset := token.NewFileSet()
	lib, err := parser.ParseFile(fset, fileName, nil, parser.Mode(0))
	if err != nil {
		return err
	}
	for _, d := range lib.Decls {
		if i, ok := d.(*ast.GenDecl); ok && i.Tok == token.IMPORT {
			*imports = append(*imports, i.Specs...)
		} else {
			*decls = append(*decls, d)
		}
	}
	return nil
}
