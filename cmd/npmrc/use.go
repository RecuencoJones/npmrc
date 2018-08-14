package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

// Use given profile
func Use(profile string) {
	// check profile is valid
	if !ProfileExists(profile) {
		fmt.Println("Profile \"" + profile + "\" does not exist")
		os.Exit(1)
	}

	// copy from $npmrc_dir/.npmrc.$profile to .npmrc
	source := path.Join(Dir, NPMRC_FILE+"."+profile)
	dest := path.Join(Home, NPMRC_FILE)

	err := cp(source, dest)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// set .npmrc_current to $profile
	// TODO

	fmt.Println("Now using profile: " + profile)
}

func cp(sourcefile, destfile string) error {
	source, err := os.Open(sourcefile)
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	dest, err := os.OpenFile(destfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	_, err = io.Copy(dest, source)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
