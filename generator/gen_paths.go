package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
	"github.com/getkin/kin-openapi/openapi3"
)

func generatePaths(f *File, t *openapi3.T, foreignPackage bool) {

	var prototype = "PrototypeEndpoint"
	if foreignPackage {
		prototype = "gac.PrototypeEndpoint"
	}

	f.Comment(GenerationTag(t))
	f.Line().Line()

	for path, pathInfo := range t.Paths {

		endpointName := CamelCase(Sluggify(path))

		f.Comment(fmt.Sprintf("#######################\n%s\n#######################\n", endpointName))
		f.Line().Line()

		for method, op := range pathInfo.Operations() {

			structName := CamelCase(fmt.Sprintf("%s_%s", method, endpointName))
			funcName := CamelCase(fmt.Sprintf("Init_%s_%s", method, endpointName))
			f.Comment(fmt.Sprintf("%s representation: %s", structName, op.Description))
			f.Type().Id(structName).Struct(
				Id(prototype),
			)

			f.Comment(fmt.Sprintf("%s is the initialization for %s", funcName, structName))
			f.Func().Id(funcName).Params().Id("Endpoint").Block(
				Id("ep").Op(":=").Op("&").Id(structName).Values(Dict{
					Id("Endpoint"): Lit(path),
					Id("Method"):   Lit(method),
					Id("Header"):   extractParamsAsDict("header", op.Parameters),
				}),
				Return(Id("ep")),
			)
		}
	}
}
