package view

import (
	"fmt"
	"io"
)

type Buf string

func Sprintf(f string, args ...interface{}) Buf {
	return Buf(fmt.Sprintf(f, args...))
}

func (b Buf) Content() (io.ReadCloser, error) { return b, nil }

func (b Buf) Read(into []byte) (int, error) {
	n := copy(into, string(b))
	return n, io.EOF
}

func (b Buf) Close() error { return nil }
