package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// ProfileExists returns whether given profile exists
func ProfileExists(profile string) bool {
	_, err := os.Stat(path.Join(Dir, NpmrcFile+"."+profile))

	return !os.IsNotExist(err)
}

// GetProfiles returns the list of available profiles
func GetProfiles() []string {
	files := LS(path.Join(Dir))

	profiles := filter(files, func(file os.FileInfo) bool {
		return strings.HasPrefix(file.Name(), ".npmrc.")
	})

	profileNames := collect(profiles, func(file os.FileInfo) string {
		return strings.Replace(file.Name(), ".npmrc.", "", 1)
	})

	return profileNames
}

// AskConfirmation via prompt with given message
func AskConfirmation(message string) bool {
	var answer string

	fmt.Printf("%s (y/N): ", message)

	_, err := fmt.Scanln(&answer)

	if err != nil {
		panic(err)
	}

	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	return answer == "y" || answer == "yes"
}

// GetEnv attempts to retrieve an env variable or returns given default value
func GetEnv(name, defaultValue string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}

	return defaultValue
}

// CP copy from source file to dest file
func CP(sourcefile, destfile string) error {
	source, err := os.Open(sourcefile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer source.Close()

	dest, err := os.OpenFile(destfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer dest.Close()

	_, err = io.Copy(dest, source)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return err
}

// LS list files in directory
func LS(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return files
}

// RM remove a file
func RM(file string) error {
	return os.Remove(file)
}

func filter(vs []os.FileInfo, f func(os.FileInfo) bool) []os.FileInfo {
	vsf := make([]os.FileInfo, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func collect(vs []os.FileInfo, f func(os.FileInfo) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
