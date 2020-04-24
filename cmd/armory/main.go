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

type ds struct {
	name string
	file string
}

var commands = map[string]ds{
	"dll":   {name: "DLL", file: "ll/dll.go"},
	"queue": {name: "Queue", file: "queue/queue.go"},
	"set":   {name: "Set", file: "set/set.go"},
	"stack": {name: "Stack", file: "stack/stack.go"},
}

func init() {
	flag.Usage = usage

	const (
		varTypeDefault = "int"
		pkgNameDefault = "main"
		outFileDefault = ""
		zeroValDefault = ""
	)

	flag.StringVar(&varType, "type", varTypeDefault, "")
	flag.StringVar(&varType, "t", varTypeDefault, "")
	flag.StringVar(&pkgName, "pkg", pkgNameDefault, "")
	flag.StringVar(&pkgName, "p", pkgNameDefault, "")
	flag.StringVar(&outFile, "out", outFileDefault, "")
	flag.StringVar(&outFile, "o", outFileDefault, "")
	flag.StringVar(&zeroVal, "zero", zeroValDefault, "")
	flag.StringVar(&zeroVal, "z", zeroValDefault, "")
}

func usage() {
	u := `  NAME:
    armory - A CLI data structure generator

  USAGE:
    armory [options] command

  COMMANDS:
    dll
    set
    sll
    stack
    queue

  OPTIONS:
    --out string, -o string     the file to write to
    --pkg string, -p string     the file's package (default "main")
    --type string, -t string    the data structure's type (default "int")
    --zero string, -z string    the type's zero value, inferred by default
    --help, -h                  display this help message
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

	c, ok := commands[flag.Arg(0)]
	if !ok {
		fmt.Printf("invalid command: %s\n", flag.Arg(0))
		flag.Usage()
		os.Exit(1)
	}

	dataStructure = c.name
	parse(c.file)
}
