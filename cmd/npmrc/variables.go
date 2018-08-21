package main

import (
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

// NpmrcFile base name
const (
	Version   = "0.1.3"
	NpmrcFile = ".npmrc"
)

var (
	// Home user home directory
	Home, _ = homedir.Dir()

	// Cwd current working directory
	Cwd, _ = os.Getwd()

	// Dir directory to store npmrc profiles
	Dir = GetEnv("NPMRC_DIR", Home)

	// Editor to use for create/update profiles
	Editor = GetEnv("EDITOR", "vi")

	// Viewer to use for displaying profiles
	Viewer = GetEnv("VIEWER", "cat")
)
