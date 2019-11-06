package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

// LowercaseReader struct
type LowercaseReader struct {
	text string
}

// NewLowercaseReader constructor
func NewLowercaseReader(text string) *LowercaseReader {
	return &LowercaseReader{text: text}
}

func (r *LowercaseReader) Read(p []byte) (int, error) {

	buf := make([]byte, len(r.text))

	for i := 0; i < len(buf); i++ {
		buf[i] = r.text[i] | 0x20
	}

	n := copy(p, buf)

	return n, io.EOF
}

func main() {
	r := NewLowercaseReader("ALL CAPITALS")
	resp, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("error: somethint went wrong when converting to lowercase")
	}

	fmt.Println(string(resp))
}
