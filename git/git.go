package main

import (
	"flag"
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	statusCode = 0
)

func main() {
	defer func() {
		os.Exit(statusCode)
	}()

	flag.Parse()

	// Check
	switch flag.Arg(0) {
	case "add":
		err := checkAdd(flag.Args()[1:])
		if err != nil {
			printErr(err)
			statusCode = 1
			return
		}
	}

	err := os.Setenv("PATH", clearPath(os.Getenv("PATH"), gopaths()))
	if err != nil {
		printErr(err)
		statusCode = 1
		return
	}

	// Exec git
	cmd := exec.Command("git", flag.Args()...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		printErr(err)
		statusCode = 1
		return
	}
}

// Remove $GOPATH/bin from $PATH
func clearPath(path string, gopath []string) string {
	paths := []string{}

SearchLoop:
	for _, p := range filepath.SplitList(path) {
		for _, gopath := range gopath {
			if p == filepath.Join(gopath, "bin") {
				continue SearchLoop
			}
		}

		paths = append(paths, p)
	}

	return strings.Join(paths, ":")
}

func printErr(err error) {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
