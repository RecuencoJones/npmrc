package main

import (
	"flag"
	"os"
)

func main() {
	useCmd := flag.NewFlagSet("use", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	if len(os.Args) < 2 {
		Help()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "use":
		useCmd.Parse(os.Args[2:])
	case "list":
		listCmd.Parse(os.Args[2:])
	case "help":
		Help()
	default:
		Help()
		os.Exit(1)
	}

	if useCmd.Parsed() {
		Use(os.Args[2])
	}

	if listCmd.Parsed() {
		List(false)
	}
}
