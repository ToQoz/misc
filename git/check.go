package main

import (
	"errors"
	"regexp"
)

func checkAdd(args []string) error {
	if contains(args, ".") {
		return errors.New("git add . is not allowed")
	}

	if contains(args, "-A") {
		pathspecs := rejectr(args, regexp.MustCompile(`^-`))

		if len(pathspecs) == 0 {
			return errors.New("git add -A without pathspec is not allowed. you should add pathspec")
		}
	}

	return nil
}
