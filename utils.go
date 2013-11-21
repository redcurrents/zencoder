package zencoder

import (
	"io"
)

type ByteReaderCloser struct {
	s []byte
	i int // current reading index
}

func NewByteReaderCloser(b []byte) *ByteReaderCloser {
	return &ByteReaderCloser{s: b}
}

func (r *ByteReaderCloser) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += n
	return
}

func (*ByteReaderCloser) Close() error {
	return nil
}
