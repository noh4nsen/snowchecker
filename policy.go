package main

import (
	"fmt"
	"os"
	"strings"
)

func runPolicies(configs []Config) {
	var errs int
	for _, config := range configs {
		var configErrs []error
		configErrs = append(configErrs, policiesDataMarts(config.DataMart)...)
		configErrs = append(configErrs, policiesDataSources(config.DataSource)...)
		configErrs = append(configErrs, policiesTeams(config.Team)...)
		logErrors(configErrs, config.filepath)
		errs += len(configErrs)
	}

	if errs > 0 {
		fmt.Printf("ERROR: %d Policies failed", errs)
		os.Exit(1)
	}
}

func logErrors(errs []error, filepath string) {
	if len(errs) > 0 {
		fmt.Printf("%d Violations found in %s file\n", len(errs), filepath)
		for index, err := range errs {
			fmt.Printf("		[%d] %s", index+1, err)
		}
	}
}

func policiesDataMarts(dataMarts map[string]DataMartDetail) []error {
	var errs []error
	for key, value := range dataMarts {
		if validateNamePattern(key, value.DataMartDomain, value.DataMartName) {
			errs = append(errs, fmt.Errorf("Invalid name pattern for %s data mart\n", key))
		}
	}
	return errs
}

func policiesDataSources(dataSources map[string]DataSourceDetail) []error {
	var errs []error
	for key, value := range dataSources {
		if exceptions(value.DataSourceDomain, "AUDM") {
			continue
		}
		if validateNamePattern(key, value.DataSourceDomain, value.DataSourceName) {
			errs = append(errs, fmt.Errorf("Invalid name pattern for %s data source\n", key))
		}
	}
	return errs
}

func validateNamePattern(resourceName string, domain string, name string) bool {
	return resourceName != domain+"__"+name
}

func exceptions(domain string, exception string) bool {
	return strings.Contains(domain, exception)
}

func policiesTeams(teams map[string][]string) []error {
	var err []error

	for key, value := range teams {
		check, wid := validateDuplicateWid(value)
		if check {
			err = append(err, fmt.Errorf("Duplicate WID %s found in team %s\n", wid, key))
		}
	}
	return err
}

func validateDuplicateWid(wids []string) (bool, string) {
	seen := make(map[string]bool)
	for _, wid := range wids {
		if seen[wid] {
			return true, wid
		}
		seen[wid] = true
	}
	return false, ""
}
