package main

import (
	"os"
)

func main() {
	validate(os.Args)
	runPolicies(parse(os.Args))
}
