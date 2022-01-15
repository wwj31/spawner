package main

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	genFile()
}
func genFile() {
	var (
		packName string
		structs  []string
	)

	pack := parsePackage(nil, nil)
	for _, astFile := range pack.Syntax {
		for name, obj := range astFile.Scope.Objects {
			if spec, ok := obj.Decl.(*ast.TypeSpec); ok {
				if _, ok2 := spec.Type.(*ast.StructType); ok2 {
					packName = astFile.Name.Name
					structs = append(structs, name)
				}
			}
		}
	}
	outputName := filepath.Join("./", strings.ToLower(fmt.Sprintf("factory.go")))
	outfile(outputName, packName, structs)
}

// parsePackage analyzes the single package constructed from the patterns and tags.
func parsePackage(patterns []string, tags []string) *packages.Package {
	cfg := &packages.Config{
		Mode:       packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	return pkgs[0]
}

const tplpackage = `{comment}
package {package}

type factory func() interface{}

func Spawner(name string) (interface{},bool) {
	f ,ok := spawner[name]
	if !ok{
		return nil,ok
	}
	return f(),true
}

var spawner = map[string]factory{
`

const field = `"{name}":func() interface{} { return &{name}{} },
`

func outfile(output, packagename string, structs []string) {
	_ = os.Remove(output)
	file, err := os.Create(output)
	if err != nil {
		log.Fatalf("error: %v outfile", err)
	}
	context := strings.NewReplacer([]string{
		"{comment}", fmt.Sprintf("// Code generated by \"spawner %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " ")),
		"{package}", packagename,
	}...).Replace(tplpackage)

	for _, val := range structs {
		r := rune(val[0])
		if r < 97 {
			continue
		}

		// provide factory func
		context += strings.NewReplacer([]string{
			"{name}", val,
		}...).Replace(field)
	}
	context += "\n}"
	_, _ = file.WriteString(context)
	_ = file.Close()
	exec.Command("go", "fmt").Output()
}

func strFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}
