package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

// EditOptions containing flag values
type EditOptions struct {
	verbose bool
	help    bool
}

// EditHandler used to parse args and options
func EditHandler(args []string, options EditOptions) {
	profile := args[0]

	if options.verbose {
		if !ProfileExists(profile) {
			fmt.Println("Creating profile \"" + profile + "\"")
		}
	}

	err := Edit(profile)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if options.verbose {
		if ProfileExists(profile) {
			fmt.Println("Saved profile \"" + profile + "\"")
		} else {
			fmt.Println("Profile \"" + profile + "\" not created")
		}
	}
}

// Edit or create given profile
func Edit(profile string) error {
	cmd := exec.Command(Editor, path.Join(Dir, NpmrcFile+"."+profile))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

// EditHelp display usage of Edit command
func EditHelp() {
	fmt.Println(`
Usage: npmrc edit [flags] <profile>

Alias: ed

Available flags:

verbose    Display additional output
h          Display this message`)
}
