package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	inp := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	out := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	totalRead := 0
	for {
		read, err := rot.r.Read(b)
		cur := totalRead
		totalRead += read
		if err == nil {
			curView := b[cur:totalRead]
			for i, v := range curView {
				idx := strings.IndexByte(inp, v)
				if idx != -1 {
					curView[i] = out[idx]
				}
			}
		}
		if err == io.EOF {
			return totalRead, err
		}
	}
}

// Rot13Demo demo
func Rot13Demo() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
