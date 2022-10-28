package goTour04

import (
	"golang.org/x/tour/reader"

	"io"
	"os"
	"strings"
)

// The io.Reader interface has a Read method, which represents the read end of a stream of data:
// func (T) Read(b []byte) (n int, err error)
// Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.

// The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.
type MyReader struct{}

func (r MyReader) Read(p []byte) (n int, err error) {
	for idx := range p {
		p[idx] = byte('A')
	}

	return len(p), nil
}

type Rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		b = 'A' + (b-'A'+13)%26
	} else if b >= 'a' && b <= 'z' {
		b = 'a' + (b-'a'+13)%26
	}
	return b
}

func (rot *Rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for idx, byteValue := range p {
		p[idx] = rot13(byteValue)
	}

	return
}

func RunExercises22_23() {
	// Reader example
	reader.Validate(MyReader{})

	// rot13Reader example
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
