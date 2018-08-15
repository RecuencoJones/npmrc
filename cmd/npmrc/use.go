package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

// UseOptions containing flag values
type UseOptions struct {
	help bool
}

// UseHandler used to parse args and options
func UseHandler(args []string, options UseOptions) {
	if options.help {
		UseHelp()
		os.Exit(0)
	}

	if len(args) < 1 {
		fmt.Println("Error: You must specify a profile!")
		UseHelp()
		os.Exit(1)
	}

	profile := args[0]

	err := Use(profile)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// Use given profile
func Use(profile string) error {
	if !ProfileExists(profile) {
		fmt.Println("Profile \"" + profile + "\" does not exist")
		os.Exit(1)
	}

	// copy from $npmrc_dir/.npmrc.$profile to .npmrc
	source := path.Join(Dir, NpmrcFile+"."+profile)
	dest := path.Join(Home, NpmrcFile)

	err := CP(source, dest)

	if err != nil {
		return err
	}

	// TODO set .npmrc_current to $profile

	fmt.Println("Now using profile: " + profile)

	return err
}

// UseHelp display usage of Use command
func UseHelp() {
	fmt.Println(`
Usage: npmrc use [flags] <profile>

Alias: u

Available flags:

h          Display this message`)
}
