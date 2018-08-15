package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

// ViewOptions containing flag values
type ViewOptions struct {
	help bool
}

// ViewHandler used to parse args and options
func ViewHandler(args []string, options ViewOptions) {
	var err error

	if options.help {
		ViewHelp()
		os.Exit(0)
	}

	if len(args) == 0 {
		err = ViewCurrent()
	} else {
		profile := args[0]

		if !ProfileExists(profile) {
			fmt.Println("Profile \"" + profile + "\" does not exist")
			os.Exit(1)
		}

		err = View(profile)
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// ViewCurrent profile
func ViewCurrent() error {
	file := path.Join(Home, NpmrcFile)

	return view(file)
}

// View given profile
func View(profile string) error {
	file := path.Join(Dir, NpmrcFile+"."+profile)

	return view(file)
}

func view(file string) error {
	cmd := exec.Command(Viewer, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

// ViewHelp display usage of View command
func ViewHelp() {
	fmt.Println(`
Usage: npmrc view [flags] <profile>

Alias: v

Available flags:

h          Display this message`)
}
