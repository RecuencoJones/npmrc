package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

// CopyOptions containing flag values
type CopyOptions struct {
	verbose bool
	help    bool
}

// CopyHandler used to parse args and options
func CopyHandler(args []string, options CopyOptions) {
	if options.help {
		CopyHelp()
		os.Exit(0)
	}

	if len(args) < 2 {
		fmt.Println("Error: You must specify source and destination profiles!")
		CopyHelp()
		os.Exit(1)
	}

	src := args[0]
	dest := args[1]

	err := Copy(src, dest)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if options.verbose {
		fmt.Println("Copied profile \"" + src + "\" to \"" + dest + "\"")
	}
}

// Copy source profile to destination profile
func Copy(srcProfile, destProfile string) error {
	var dest string

	// check profile is valid
	if !ProfileExists(srcProfile) {
		fmt.Println("Error: Profile \"" + srcProfile + "\" does not exist")
		os.Exit(1)
	}

	if destProfile == "." {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		dest = path.Join(dir, NpmrcFile)
	} else {
		ValidateProfile(destProfile)
		dest = path.Join(Dir, NpmrcFile+"."+destProfile)
	}

	// copy from $npmrc_dir/.npmrc.$srcProfile to $npmrc_dir/.npmrc.$destProfile
	source := path.Join(Dir, NpmrcFile+"."+srcProfile)

	return CP(source, dest)
}

// CopyHelp display usage of Copy command
func CopyHelp() {
	fmt.Println(`
Usage: npmrc copy [flags] <sourceProfile> <destProfile>

Alias: cp

Available flags:

verbose    Display additional output
h          Display this message

Details:

If <destProfile> is ".", the selected profile is copied to current working directory.`)
}
