package main

import (
	"fmt"
	"os"
	"strings"
)

// ListOptions containing flag values
type ListOptions struct {
	asList bool
	help   bool
}

// ListHandler used to parse args and options
func ListHandler(options ListOptions) {
	if options.help {
		ListHelp()
		os.Exit(0)
	}

	List(options.asList)
}

// List available profiles
func List(asList bool) {
	profileNames := GetProfiles()

	if len(profileNames) == 0 {
		fmt.Println("No profiles available")
	} else {
		if asList {
			for _, profileName := range profileNames {
				fmt.Println(profileName)
			}
		} else {
			fmt.Println(strings.Join(profileNames, " "))
		}
	}
}

// ListHelp display usage of List command
func ListHelp() {
	fmt.Println(`
Usage: npmrc list [flags]

Alias: ls

Available flags:

l          Display as list
h          Display this message`)
}
