package main

import (
	"testing"
)

func TestClearPath(t *testing.T) {
	var expected string
	var got string

	expected = "/usr/bin"
	got = clearPath("/usr/bin:/a/go/path/bin", []string{"/a/go/path"})
	if expected != got {
		t.Errorf("expected=%s, but got=%s", expected, got)
	}

	expected = "/usr/bin"
	got = clearPath("/usr/bin:/a/go/path/bin:/b/go/path/bin", []string{"/a/go/path", "/b/go/path"})
	if expected != got {
		t.Errorf("expected=%s, but got=%s", expected, got)
	}
}
