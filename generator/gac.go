package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	DoDebug         bool
	InputFile       string
	OutputDir       string
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
	flag.BoolVar(&CopyBoilerplate, "include-boilerplate", true, "also generate boilerplate code")
	flag.BoolVar(&CopyRuntime, "include-runtime", true, "also export necessary runtime")
	flag.BoolVar(&OutputVersion, "v", false, "output version and exit")

	flag.Parse()
	InputFile = flag.Arg(0)

	if DoDebug {
		fmt.Printf("DoDebug: %v\n", DoDebug)
		fmt.Printf("InputFile: %s\n", InputFile)
		fmt.Printf("OutputDir: %s\n", OutputDir)
		fmt.Printf("CopyBoilerplate: %v\n", CopyBoilerplate)
		fmt.Printf("CopyRuntime: %v\n", CopyRuntime)
		fmt.Printf("OutputVersion: %v\n", OutputVersion)
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

	fmt.Printf("%#v\n", oa3)
}
