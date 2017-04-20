package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(c byte) byte {

	// HACK:ifとswitch迷いました
	switch {
	case 'A' <= c && c <= 'Z':
		return (c-'A'+13)%26 + 'A'
	case 'a' <= c && c <= 'z':
		return (c-'a'+13)%26 + 'a'
	}
	return c
}

func (rot13r *rot13Reader) Read(p []byte) (int, error) {
	n, err := rot13r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
