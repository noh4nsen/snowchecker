package main

import (
	"fmt"
	"os"
	"strings"
)

func validate(args []string) {
	errs := checkInputArgs(args)
	for _, value := range errs {
		fmt.Println(value)
	}
	if len(errs) > 0 {
		os.Exit(1)
	}
}

func checkInputArgs(args []string) []error {
	var errs []error
	if len(args) <= 1 {
		errs = append(errs, fmt.Errorf("ERROR: no tfvars files provided, at least one is required"))
	}
	errs = append(errs, checkFileFormat(args)...)
	return errs
}

func checkFileFormat(args []string) []error {
	var errs []error
	for index, value := range args {
		if index > 0 {
			if !strings.Contains(value, ".tfvars") {
				errs = append(errs, fmt.Errorf("ERROR: only .tfvars files allowed, provided %s", value))
			}
		}
	}
	return errs
}
