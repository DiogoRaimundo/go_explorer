package main

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

func (r MyReader) Read(array []byte) (int, error) {
	for idx := range array {
		array[idx] = byte('A')
	}

	return len(array), nil
}

type rot13Reader struct {
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

func (rot *rot13Reader) Read(bytes []byte) (nBytes int, err error) {
	nBytes, err = rot.r.Read(bytes)
	for idx, byteValue := range bytes {
		bytes[idx] = rot13(byteValue)
	}

	return
}

func main() {
	// Reader example
	reader.Validate(MyReader{})

	// rot13Reader example
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
