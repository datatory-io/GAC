package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/hashicorp/go-version"
	"os"
)

var (
	DoDebug         bool
	InputFile       string
	OutputDir       string
	PackageName     string
	CopyBoilerplate bool
	CopyRuntime     bool
	OutputVersion   bool
)

func main() {

	if DoDebug {
		debugInfo()
	}

	flag.BoolVar(&DoDebug, "d", true, "debug output")
	flag.StringVar(&OutputDir, "o", "", "the output directory")
	flag.StringVar(&PackageName, "p", "", "the detigned package name - leave blank for gac")
	flag.BoolVar(&CopyBoilerplate, "include-boilerplate", true, "also generate boilerplate code")
	flag.BoolVar(&CopyRuntime, "include-runtime", true, "also export necessary runtime")
	flag.BoolVar(&OutputVersion, "v", false, "output version and exit")

	flag.Parse()
	InputFile = Untrail(flag.Arg(0))
	OutputDir = Untrail(OutputDir)

	if DoDebug {
		fmt.Printf("DoDebug: %v\n", DoDebug)
		fmt.Printf("InputFile: %s\n", InputFile)
		fmt.Printf("OutputDir: %s\n", OutputDir)
		fmt.Printf("PackageName: %s\n", PackageName)
		fmt.Printf("CopyBoilerplate: %v\n", CopyBoilerplate)
		fmt.Printf("CopyRuntime: %v\n", CopyRuntime)
		fmt.Printf("OutputVersion: %v\n", OutputVersion)
	}

	if PackageName == "" {
		PackageName = "gac"
	}

	if flag.NArg() < 1 {
		displayUsage()
		os.Exit(1)
	}

	if OutputVersion {
		displayVersion()
		os.Exit(0)
	}

	if _, err := os.Stat(InputFile); err != nil {
		exitErr("unable to read file: %s (ERR: %s)", InputFile, err)
	}

	if info, err := os.Stat(OutputDir); errors.Is(err, os.ErrNotExist) || !info.IsDir() {
		if DoDebug {
			_, _ = fmt.Fprintf(os.Stderr, "%#v\n", info)
		}
		exitErr("directory %s does not exist or is not writable (ERR: %s)", OutputDir, err)
	}

	oa3, err := LoadSwagger(InputFile)
	if err != nil {
		exitErr("error loading swagger for file: %s (ERR: %s)", InputFile, err)
	}

	v1, _ := version.NewVersion("3.0.0")
	v2, err := version.NewVersion(oa3.OpenAPI)
	if err != nil {
		exitErr("Unable to understand file format of file: %s (ERR: %s)", InputFile, err)
	}

	if v1.GreaterThan(v2) {
		exitErr("Only OPENAPI 3.0.0 and greater is supported")
	}

	// ##### ENVIRONMENT

	environment := jen.NewFile(PackageName)

	generateEnvironments(environment, oa3, PackageName != "gac")

	fh, err := os.Create(OutputDir + "/environment.go")
	if err != nil {
		exitErr("unable to open file for writing in dir %s (ERR: %s)", OutputDir, err)
	}
	_, _ = fmt.Fprintf(fh, "%#v", environment)
	fh.Close()

}
