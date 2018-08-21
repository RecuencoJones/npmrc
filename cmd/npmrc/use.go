package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

// UseOptions containing flag values
type UseOptions struct {
	local bool
	help  bool
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

	err := Use(profile, options.local)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// Use given profile
func Use(profile string, local bool) error {
	var dest string

	if !ProfileExists(profile) {
		fmt.Println("Profile \"" + profile + "\" does not exist")
		os.Exit(1)
	}

	// copy from $npmrc_dir/.npmrc.$profile to .npmrc
	source := path.Join(Dir, NpmrcFile+"."+profile)

	if local {
		dest = path.Join(Cwd, NpmrcFile)
	} else {
		dest = path.Join(Home, NpmrcFile)
	}

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

local      Use the profile for current directory
h          Display this message`)
}
