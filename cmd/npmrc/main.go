package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	useCmd := flag.NewFlagSet("use", flag.ExitOnError)
	useLocalFlag := useCmd.Bool("local", false, "Use the profile for current directory")
	useHelpFlag := useCmd.Bool("h", false, "Display help")

	viewCmd := flag.NewFlagSet("view", flag.ExitOnError)
	viewHelpFlag := viewCmd.Bool("h", false, "Display help")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listAsListFlag := listCmd.Bool("l", false, "Display results as list")
	listHelpFlag := listCmd.Bool("h", false, "Display help")

	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	editVerboseFlag := editCmd.Bool("verbose", false, "Display additional output")
	editHelpFlag := editCmd.Bool("h", false, "Display help")

	copyCmd := flag.NewFlagSet("copy", flag.ExitOnError)
	copyVerboseFlag := copyCmd.Bool("verbose", false, "Display additional output")
	copyHelpFlag := copyCmd.Bool("h", false, "Display help")

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeVerboseFlag := removeCmd.Bool("verbose", false, "Display additional output")
	removeForceFlag := removeCmd.Bool("f", false, "Force removal")
	removeHelpFlag := removeCmd.Bool("h", false, "Display help")

	if len(os.Args) < 2 {
		Help()
		os.Exit(1)
	}

	args := os.Args[2:]

	switch os.Args[1] {
	case "use":
		fallthrough
	case "u":
		useCmd.Parse(args)
	case "view":
		fallthrough
	case "v":
		viewCmd.Parse(args)
	case "list":
		fallthrough
	case "ls":
		listCmd.Parse(args)
	case "copy":
		fallthrough
	case "cp":
		copyCmd.Parse(args)
	case "edit":
		fallthrough
	case "ed":
		editCmd.Parse(args)
	case "remove":
		fallthrough
	case "rm":
		removeCmd.Parse(args)
	case "help":
		fallthrough
	case "h":
		Help()
	case "version":
		fmt.Println(Version)
	default:
		Help()
		os.Exit(1)
	}

	// Use
	if useCmd.Parsed() {
		options := UseOptions{
			local: *useLocalFlag,
			help:  *useHelpFlag,
		}

		UseHandler(useCmd.Args(), options)
	}

	// View
	if viewCmd.Parsed() {
		options := ViewOptions{
			help: *viewHelpFlag,
		}

		ViewHandler(viewCmd.Args(), options)
	}

	// List
	if listCmd.Parsed() {
		options := ListOptions{
			asList: *listAsListFlag,
			help:   *listHelpFlag,
		}

		ListHandler(options)
	}

	// Copy
	if copyCmd.Parsed() {
		options := CopyOptions{
			verbose: *copyVerboseFlag,
			help:    *copyHelpFlag,
		}

		CopyHandler(copyCmd.Args(), options)
	}

	// Edit
	if editCmd.Parsed() {
		options := EditOptions{
			verbose: *editVerboseFlag,
			help:    *editHelpFlag,
		}

		EditHandler(editCmd.Args(), options)
	}

	// Remove
	if removeCmd.Parsed() {
		options := RemoveOptions{
			force:   *removeForceFlag,
			verbose: *removeVerboseFlag,
			help:    *removeHelpFlag,
		}

		RemoveHandler(removeCmd.Args(), options)
	}
}

// Help display usage of npmrc command
func Help() {
	fmt.Println(`
Usage: npmrc <command>

Available commands:

use, u      Select a profile
view, v     View a profile
list, ls    List available profiles
edit, ed    Create or update profiles
copy, cp    Copy profiles
remove, rm  Remove a profile
help, h     Display this message
version     Display version

Profile names:

A profile names must comply with the following:

- Must not start with "."
- Must not contain spaces
- Must be of at least 1 character length`)
}
