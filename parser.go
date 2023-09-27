package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

func parse(filePaths []string) []Config {
	var configs []Config
	for index, filePath := range filePaths {
		if index != 0 {
			var config Config
			config.filepath = filePath
			decode(parseFile(filePath), &config)
			configs = append(configs, config)
		}
	}
	return configs
}

func parseFile(filePath string) hcl.Body {
	parser := hclparse.NewParser()
	file, diagnostics := parser.ParseHCLFile(filePath)
	if diagnostics.HasErrors() {
		fmt.Println("ERROR: ", diagnostics.Error())
		os.Exit(1)
	}
	return file.Body
}

func decode(body hcl.Body, config *Config) {
	diagnostics := gohcl.DecodeBody(body, nil, config)
	if diagnostics.HasErrors() {
		fmt.Println("ERROR: ", diagnostics.Error())
		os.Exit(1)
	}
}
