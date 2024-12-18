package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/noperator/jqfmt"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Fatalf(message, err)
	}
}

func main() {
	// funcsStr := flag.String("fn", "", "functions")
	opsStr := flag.String("op", "", "operators")
	obj := flag.Bool("ob", false, "objects")
	arr := flag.Bool("ar", false, "arrays")
	oneLn := flag.Bool("o", false, "one line")
	file := flag.String("f", "", "file")
	verbose := flag.Bool("v", false, "verbose")
	// noHang := flag.Bool("nh", false, "no hanging indent")
	flag.Parse()
	var from_stdin bool = false

	if !*verbose {
		log.SetOutput(io.Discard)
	}

	if *file == "" {
		from_stdin = true
	}

	// var funcs []string
	// if *funcsStr == "" {
	// 	funcs = []string{}
	// } else {
	// 	funcs = strings.Split(*funcsStr, ",")
	// }

	var ops []string
	if *opsStr == "" {
		ops = []string{}
	} else {
		ops = strings.Split(*opsStr, ",")
	}

	cliCfg, err := jqfmt.ValidateConfig(jqfmt.JqFmtCfg{
		Arr: *arr,
		// Funcs: funcs,
		Obj:   *obj,
		OneLn: *oneLn,
		Ops:   ops,
		// Hang:  !(*noHang),
	})
	assertErrorToNilf("invalid config: %v", err)

	var jqBytes []byte
	if from_stdin {
		jqBytes, err = io.ReadAll(os.Stdin)
	} else {
		jqBytes, err = os.ReadFile(*file)
	}
	assertErrorToNilf("could not read file: %v", err)
	jqStr := string(jqBytes)
	jqStrFmt, err := jqfmt.DoThing(jqStr, cliCfg)
	assertErrorToNilf("could not format jq: %v", err)
	fmt.Print(jqStrFmt)
}
