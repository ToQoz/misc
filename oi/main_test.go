package main

import (
	"os"
	"regexp"
	"runtime"
	"testing"
)

func TestGoi(t *testing.T) {
	if envvar := os.Getenv("GOMAXPROCS"); envvar == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	for i := 0; i < 500*runtime.NumCPU(); i++ {
		go func() {
			got := goi()
			pattern := `(?:オー{0,10}イ){1,30}オーーーーーーーーーーイ!!!!!!!!!!`
			if !regexp.MustCompile(pattern).Match(got) {
				t.Error(string(got) + " does not match to " + pattern)
			}
		}()
	}
}
