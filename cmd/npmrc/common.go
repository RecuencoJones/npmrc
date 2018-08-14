package main

import (
	"os"
	"path"
)

// ProfileExists returns whether given profile exists
func ProfileExists(profile string) bool {
	_, err := os.Stat(path.Join(Dir, NPMRC_FILE+"."+profile))

	return !os.IsNotExist(err)
}
