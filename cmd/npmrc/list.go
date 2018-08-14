package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func List(asList bool) {
	profileNames := getProfiles()

	if asList {
		for _, profileName := range profileNames {
			fmt.Println(profileName)
		}
	} else {
		fmt.Println(strings.Join(profileNames, ", "))
	}
}

func getProfiles() []string {
	files := ls(path.Join(Dir))

	profiles := filter(files, func(file os.FileInfo) bool {
		return strings.HasPrefix(file.Name(), ".npmrc.")
	})

	profileNames := collect(profiles, func(file os.FileInfo) string {
		return strings.Replace(file.Name(), ".npmrc.", "", 1)
	})

	return profileNames
}

func ls(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return files
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
