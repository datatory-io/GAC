package main

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gosimple/slug"
	"github.com/iancoleman/strcase"
	"log"
	"net/url"
	"os"
	"runtime/debug"
	"strings"
)

func exitErr(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func debugInfo() {
	fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))

	cwd, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("  cwd = %s\n", cwd)
	fmt.Printf("  os.Args = %#v\n", os.Args)
}

func displayUsage() {
	_, _ = fmt.Fprintln(os.Stderr, "GAC Generator ", getVersion())
	_, _ = fmt.Fprintln(os.Stderr, "USAGE: go generate -o /path/to/output /path/to/openapi.yaml")
}

func displayVersion() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		_, _ = fmt.Fprintln(os.Stderr, "error reading build info")
		os.Exit(1)
	}
	fmt.Println(bi.Main.Path + "/exec/gac_generate")
	fmt.Println(bi.Main.Version)
}

func Untrail(in string) string {
	return strings.TrimSuffix(in, "/")
}

func getVersion() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		_, _ = fmt.Fprintln(os.Stderr, "error reading build info")
		os.Exit(1)
	}
	return fmt.Sprintf("%s (GOLANG version %s)", bi.Main.Version, bi.GoVersion)
}

func LoadSwagger(filePath string) (swagger *openapi3.T, err error) {

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	u, err := url.Parse(filePath)
	if err == nil && u.Scheme != "" && u.Host != "" {
		return loader.LoadFromURI(u)
	} else {
		return loader.LoadFromFile(filePath)
	}
}

func Sluggify(in string) string {
	return slug.Make(in)
}

func CamelCase(in string) string {
	return strcase.ToCamel(in)
}

func AssembleServerUri(url string, vars map[string]*openapi3.ServerVariable) string {

	if len(vars) < 1 {
		return url
	}

	uri := url
	for name, svar := range vars {
		uri = strings.ReplaceAll(uri, "{"+name+"}", svar.Default)
	}

	return uri
}
