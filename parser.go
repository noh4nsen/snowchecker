package main

import (
	"log"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

func parse(filePath string) Config {
	var config Config

	decode(parseFile(filePath), &config)

	return config
}

func parseFile(filePath string) hcl.Body {
	parser := hclparse.NewParser()

	file, diagnostics := parser.ParseHCLFile(filePath)

	if diagnostics.HasErrors() {
		log.Fatal(diagnostics.Error())
	}

	return file.Body
}

func decode(body hcl.Body, config *Config) {
	diagnostics := gohcl.DecodeBody(body, nil, config)

	if diagnostics.HasErrors() {
		log.Fatal(diagnostics.Error())
	}
}
