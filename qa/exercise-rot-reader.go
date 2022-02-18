package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (v rot13Reader) Read(b []byte) (int, error) {
	n, err := v.r.Read(b)
	for i := 0; i < n ; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func rot13(b byte) byte {
	//ASCII 65 is A and 90 is Z
	if b > 64 && b < 91 {
		b = ((b - 65 + 13) % 26) + 65
	}
	//ASCII 97 is a and 122 is z
	if b > 96 && b < 123 {
		b = ((b - 97 + 13) % 26) + 97
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
