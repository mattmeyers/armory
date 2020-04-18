package main

import (
	"flag"
	"os"
	"strings"
)

var dataStructure string
var varType string
var varTypeName string
var pkgName string
var outFile string

func init() {
	flag.StringVar(&varType, "type", "int", "the data type for the data structure")
	flag.StringVar(&pkgName, "pkg", "main", "the package that the file will belong to")
	flag.StringVar(&outFile, "out", "", "the file to write to")
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
		parse("set.go")
	default:
		flag.Usage()
		os.Exit(1)
	}

}
