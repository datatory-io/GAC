package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/getkin/kin-openapi/openapi3"
)

func generateEnvironments(f *File, t *openapi3.T, foreignPackage bool) {

	var prototype = "PrototypeEnvironment"
	if foreignPackage {
		prototype = "gac.PrototypeEnvironment"
	}

	for idx, srv := range t.Servers {

		ident := fmt.Sprintf("srv%d", idx)
		name := CamelCase(Sluggify(srv.Description))
		if name == "" {
			name = CamelCase(Sluggify(srv.URL))
		}

		f.Type().Id(name).Struct(
			Id(prototype),
			Id("Name").String(),
			Id("Url").String(),
		)

		url := AssembleServerUri(srv.URL, srv.Variables)

		f.Func().Id("Init"+name).Params().Id("Environment").Block(
			Id("op").Op(":=").Op("&").Id(name).Values(Dict{
				Id("Name"): Lit(name),
				Id("Url"):  Lit(url),
			}),
			Return(Id("op")),
		)

		f.Func().Params(Id(ident).Op("*").Id(name)).Id("GetName").Params().String().Block(
			Return(Id(ident).Dot("Name")),
		)

		f.Func().Params(Id(ident).Op("*").Id(name)).Id("GetUri").Params().String().Block(
			Return(Id(ident).Dot("Url")),
		)
	}
}
