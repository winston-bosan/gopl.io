package main

import "io"

type LReader struct {
	str string
}

func (r *LReader) Read(p []byte) (n int, err error) {
	temp_n := copy(p, r.str)
	if temp_n >= n {
		err = io.EOF
	}
	return 
}

func limitReader(r io.Reader, n int) io.Reader {

}

func main() {

}