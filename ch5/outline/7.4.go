package main

import (
	"io"
)

type NewReader struct {
	str string
}

func (r *NewReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.str)
	r.str = r.str[n:]
	if len(r.str) == 0 {
		err = io.EOF
	}
	return n, err
}

func NeuReader(s string) io.Reader {
	return &NewReader{s}
}