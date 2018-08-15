package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

// RemoveOptions containing flag values
type RemoveOptions struct {
	force   bool
	verbose bool
	help    bool
}

// RemoveHandler used to parse args and options
func RemoveHandler(args []string, options RemoveOptions) {
	if options.help {
		RemoveHelp()
		os.Exit(0)
	}

	if len(args) < 1 {
		fmt.Println("Error: You must specify a profile!")
		RemoveHelp()
		os.Exit(1)
	}

	if !options.force {
		// TODO prompt
	}

	profile := args[0]

	if !ProfileExists(profile) {
		if options.verbose {
			fmt.Println("Profile \"" + profile + "\" does not exist. Nothing to do here.")
		}

		os.Exit(0)
	}

	err := Remove(profile)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if options.verbose {
		fmt.Println("Removed profile \"" + profile + "\"")
	}
}

// Remove given profile
func Remove(profile string) error {
	return RM(path.Join(Dir, NpmrcFile+"."+profile))
}

// RemoveHelp display usage of Remove command
func RemoveHelp() {
	fmt.Println(`
Usage: npmrc remove [flags] <profile>

Alias: rm

Available flags:

verbose    Display additional output
f          Force removal
h          Display this message`)
}
