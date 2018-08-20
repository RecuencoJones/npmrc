package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// IsProfileValid checks if a given profile name complies with the naming rules
func IsProfileValid(profile string) (bool, []error) {
	var errs []error

	if strings.HasPrefix(profile, ".") {
		err := errors.New("Profile name cannot start with \".\"")

		errs = append(errs, err)
	}

	if strings.Contains(profile, " ") {
		err := errors.New("Profile name cannot contain spaces")

		errs = append(errs, err)
	}

	if len(profile) < 1 {
		err := errors.New("Profile name must be at least one character long")

		errs = append(errs, err)
	}

	return len(errs) == 0, errs
}

// ValidateProfile aborts execution if profile is non compliant
func ValidateProfile(profile string) {
	isValid, errors := IsProfileValid(profile)

	if !isValid {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
