package main

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/exp/slices"
)

func main() {
	config := parse("./mock/test.tfvars")

	var wids []string

	for key, value := range config.Team {
		if key == "PENG" {
			fmt.Println("PENG: ", value)
		}
		for _, v := range value {
			if !slices.Contains(wids, v) {
				wids = append(wids, v)
			}
		}
	}

	fmt.Println()
	fmt.Println("Total users: ", len(wids))
	fmt.Println()

	for key, value := range config.DataSource {
		if key != value.DataSourceDomain+"__"+value.DataSourceName {
			fmt.Println("- MART: ", key)
			fmt.Println("-- VALUE: ", value.DataSourceDomain+"__"+value.DataSourceName)
			fmt.Println(false)
		}
	}
	fmt.Println()

	for key, value := range config.DataMart {
		if key != value.DataMartDomain+"__"+value.DataMartName {
			fmt.Println("- MART: ", key)
			fmt.Println("-- VALUE: ", value.DataMartDomain+"__"+value.DataMartName)
			fmt.Println(false)
		}
	}
	fmt.Println()

	output, err := json.MarshalIndent(config, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
