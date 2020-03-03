package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

type SaveOptions struct {
	verbose bool
	force   bool
	help    bool
}

func SaveHandler(args []string, options SaveOptions) {
	if options.help {
		SaveHelp()
		os.Exit(0)
	}

	if len(args) < 1 {
		fmt.Println("Error: You must specify destination profile!")
		SaveHelp()
		os.Exit(1)
	}

	profile := args[0]

	err := Save(profile)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if options.verbose {
		fmt.Println("Saved current global profile to \"" + profile + "\"")
	}
}

func Save(profile string) {
	ValidateProfile(profile)

	if ProfileExists(profile) {
		// prompt
	}

	source := path.Join(Dir, NpmrcFile)
	dest := path.Join(Dir, NpmrcFile+"."+profile)

	return CP(source, dest)
}

// SaveHelp display usage of Save command
func SaveHelp() {
	fmt.Println(`
Usage: npmrc save [flags] <profile>

Alias: sv

Available flags:

verbose    Display additional output
force      Skip prompt
h          Display this message

Details:

Saved global profile to given <profile> name.`)
}
