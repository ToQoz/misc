package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	n = flag.Int("n", 1, "number of lines")
)

func main() {
	flag.Parse()

	if *n <= 0 {
		fmt.Fprintln(os.Stderr, "-n should be larger than 0")
	}

	for i := 0; i < *n; i++ {
		os.Stdout.Write(goi())
		os.Stdout.Write([]byte("\n"))
	}
}

func goi() []byte {
	buf := make([]byte, 0, 200)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// (?:オー{0, 10}イ){1, 30}オーーーーーーーーーーイ!!!!!!!!!!
	for i, c := 0, r.Int()%29+1; i < c; i++ {
		buf = append(buf, []byte("オ")...)

		if r.Int()%3 == 0 {
			for i, c := 0, r.Int()%10; i < c; i++ {
				buf = append(buf, []byte("ー")...)
			}
		}

		buf = append(buf, []byte("イ")...)
	}
	buf = append(buf, []byte("オーーーーーーーーーーイ!!!!!!!!!!")...)

	return buf
}
