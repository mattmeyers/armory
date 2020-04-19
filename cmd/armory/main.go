package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var dataStructure string
var varType string
var varTypeName string
var pkgName string
var outFile string
var zeroVal string

func init() {
	flag.Usage = usage

	const (
		varTypeDefault = "int"
		varTypeUsage   = "the data type for the data structure"
		pkgNameDefault = "main"
		pkgNameUsage   = "the package that the file will belong to"
		outFileDefault = ""
		outFileUsage   = "the file to write to"
		zeroValDefault = ""
		zeroValueUsage = "the type's zero value, inferred by default"
	)

	flag.StringVar(&varType, "type", varTypeDefault, varTypeUsage)
	flag.StringVar(&varType, "t", varTypeDefault, varTypeUsage)
	flag.StringVar(&pkgName, "pkg", pkgNameDefault, pkgNameUsage)
	flag.StringVar(&pkgName, "p", pkgNameDefault, pkgNameUsage)
	flag.StringVar(&outFile, "out", outFileDefault, outFileUsage)
	flag.StringVar(&outFile, "o", outFileDefault, outFileUsage)
	flag.StringVar(&zeroVal, "zero", zeroValDefault, zeroValueUsage)
	flag.StringVar(&zeroVal, "z", zeroValDefault, zeroValueUsage)
}

func usage() {
	u := `  armory - A CLI data structure generator
    armory [options] <data structure>

  data structures:
    set
    stack

  options:
    --out string, -o string   the file to write to
    --pkg string, -p string   the file's package (default "main")
    --type string, -t string  the data structure's type (default "int")
    --zero string, -z string  the type's zero value, inferred by default
    --help, -h                display this help message
  `
	fmt.Fprint(flag.CommandLine.Output(), u)
}

func main() {
	flag.Parse()
	varTypeName = strings.Title(varType)

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case "set":
		dataStructure = "Set"
		parse("set/set.go")
	case "stack":
		dataStructure = "Stack"
		parse("stack/stack.go")
	default:
		flag.Usage()
		os.Exit(1)
	}

}
