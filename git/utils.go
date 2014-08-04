package main

import (
	"go/build"
	"os"
	"path/filepath"
	"regexp"
)

func gopaths() []string {
	paths := filepath.SplitList(os.Getenv("GOPATH"))

	for _, p := range filepath.SplitList(build.Default.GOPATH) {
		if !contains(paths, p) {
			paths = append(paths, p)
		}
	}

	return paths
}

func contains(es []string, t string) bool {
	for _, e := range es {
		if e == t {
			return true
		}
	}

	return false
}

func selectr(es []string, r *regexp.Regexp) []string {
	ret := []string{}

	for _, e := range es {
		if r.Match([]byte(e)) {
			ret = append(ret, e)
		}
	}

	return ret
}

func rejectr(es []string, r *regexp.Regexp) []string {
	ret := []string{}

	for _, e := range es {
		if !r.Match([]byte(e)) {
			ret = append(ret, e)
		}
	}

	return ret
}
