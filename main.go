package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

var (
	dbType     = flag.String("type", "", "This is the DB type - currently we support RethinkDB")
	fileName   = flag.String("filename", "", "This is the file name to be scanned")
	outputFile = flag.String("output", "", "This is the name of the output file")

	validDBTypes = map[string]struct{}{"rethinkdb": {}}
)

func listDBtypes() string {
	output := ""
	for k := range validDBTypes {
		output += ", " + k
	}
	return output
}

func main() {
	flag.Parse()
	if *dbType == "" {
		flag.PrintDefaults()
		log.Fatalln("Error: You must choose a DB type")
	}

	if _, ok := validDBTypes[strings.ToLower(*dbType)]; !ok {
		flag.PrintDefaults()
		log.Fatalln("You must choose a valid DB type:", listDBtypes())
	}

	if *fileName == "" {
		flag.PrintDefaults()
		log.Fatalln("Error: You specify a valid file name")
	}

	if *outputFile == "" || !strings.HasSuffix(*outputFile, ".go") {
		flag.PrintDefaults()
		log.Fatalln("Error: Output file must not be empty and must have the suffix of '.go'")
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, *fileName, nil, parser.AllErrors)
	if err != nil {
		log.Fatalln("Unable to parse file submitted, please fix the errors", err)
	}

	if ok := ast.FileExports(f); !ok {
		log.Fatalln("You must have exported fields for this program to work")
	}

	switch strings.ToLower(*dbType) {
	case "rethinkdb":
		_, err := generateRethinkDBMethods(f)
		if err != nil {
			log.Fatalln("Unable to generate RethinkDB methods", err)
		}

	}
}
