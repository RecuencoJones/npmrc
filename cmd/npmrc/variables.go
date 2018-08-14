package main

import homedir "github.com/mitchellh/go-homedir"

const NPMRC_FILE = ".npmrc"

var (
	// Home user home directory
	Home, _ = homedir.Dir()

	// Dir directory to store npmrc profiles
	Dir = GetEnv("NPMRC_DIR", Home)

	// Editor to use for create/update profiles
	Editor = GetEnv("EDITOR", "vim")
)
