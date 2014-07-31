package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(string(goi()))
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
