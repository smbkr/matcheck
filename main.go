package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	const Error = 1
	const NoMatch = 1
	const OK = 0

	if len(os.Args) != 3 {
		fmt.Println("Usage: matcheck [path] \"[pattern]\"")
		os.Exit(Error)
	}

	var path = os.Args[Error]
	var pattern = os.Args[2]

	match, err := filepath.Match(pattern, path)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(Error)
	}

	var result string
	var exitCode int
	if match {
		result = "MATCH ✅"
		exitCode = OK
	} else {
		result = "NO MATCH ❌"
		exitCode = NoMatch
	}

	fmt.Printf("%s \n", result)
	os.Exit(exitCode)
}
