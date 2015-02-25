package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var exitCode = 0

func report(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	flag.Usage()
	exitCode = 2
}

var (
	// Global flags
	desc = "Perform 'ghq' using ghq.root[i]"
	i    = flag.Int("i", 0, "Specify index of ghq.root")

	// Root flags
	rootDesc = "Perform 'ghq root' using ghq.root[i]"
	rootFS   = flag.NewFlagSet("ghq-rootn root", flag.ExitOnError)

	// List flags
	listDesc      = "Perform 'ghq list' using ghq.root[i]"
	listFS        = flag.NewFlagSet("ghq-rootn list", flag.ExitOnError)
	printFullpath = listFS.Bool("p", false, "Print full-paths")
	exact         = listFS.Bool("e", false, "Perform an exact match")
)

func usage() {
	fmt.Fprintf(os.Stderr, `NAME:
  ghq-rootn - %s

USAGE:
  ghq-rootn [global options] command [command options] [arguments...]

GLOBAL OPTIONS:
`, desc)

	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")

	fmt.Fprintf(os.Stderr, `COMMANDS:
  root: %s
  list: %s

`, rootDesc, listDesc)
	os.Exit(1)
}

func usageRoot() {
	fmt.Fprintf(os.Stderr, `NAME:
  root - %s

USAGE:
  ghq-rootn [global options] root

GLOBAL OPTIONS:
`, rootDesc)

	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")

	os.Exit(1)
}

func usageList() {
	fmt.Fprintf(os.Stderr, `NAME:
  list - %s

USAGE:
  ghq-rootn [global options] list [command options] [<query>]

GLOBAL OPTIONS:
`, listDesc)

	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")

	fmt.Fprintf(os.Stderr, "COMMAND OPTIONS:\n")
	listFS.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")

	os.Exit(1)
}

func init() {
	flag.Usage = usage
	listFS.Usage = usageList
	rootFS.Usage = usageRoot
	flag.Parse()
}

func main() {
	defer func() {
		os.Exit(exitCode)
	}()

	if err := doMain(); err != nil {
		report(err)
		return
	}
}

func doMain() error {
	command := flag.Arg(0)

	root, err := Root(*i)
	if err != nil {
		return err
	}

	switch command {
	case "root":
		rootFS.Parse(flag.Args()[1:])

		fmt.Fprintln(os.Stdout, root)
	case "list":
		listFS.Parse(flag.Args()[1:])

		listArgs := []string{"-p"}
		if *exact {
			listArgs = append(listArgs, "-e")
		}
		if q := listFS.Arg(0); q != "" {
			listArgs = append(listArgs, q)
		}

		paths, err := GhqList(listArgs...)
		if err != nil {
			return err
		}

		prefix := root + string(filepath.Separator)
		for _, p := range paths {
			if !strings.HasPrefix(p, prefix) {
				continue
			}

			if *printFullpath {
				fmt.Fprintln(os.Stdout, p)
			} else {
				fmt.Fprintln(os.Stdout, strings.TrimPrefix(p, prefix))
			}
		}
	default:
		flag.Usage()
	}

	return nil
}

// Root returns ghq.root[i]
func Root(i int) (string, error) {
	roots, err := GhqRoot("-all")
	if err != nil {
		return "", err
	}

	if i >= len(roots) {
		return "", fmt.Errorf("specified index is out of ghq.root's range. [len(ghq.root)=>%d]", len(roots))
	}

	return roots[i], nil
}

// GhqList returns output of 'ghq list'
func GhqList(args ...string) ([]string, error) {
	listArgs := append([]string{"list"}, args...)
	return ghq(listArgs)
}

// GhqRoot returns output of 'ghq root'
func GhqRoot(args ...string) ([]string, error) {
	rootArgs := append([]string{"root"}, args...)
	return ghq(rootArgs)
}

func ghq(args []string) ([]string, error) {
	cmd := exec.Command("ghq", args...)
	cmd.Stderr = os.Stderr
	buf, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimRight(string(buf), "\n"), "\n"), nil
}
