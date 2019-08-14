package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (n int, err error) {
	n, err = reader.r.Read(b)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		b[i] = modifyToRot13(b[i])
	}
	return n, err
}

func modifyToRot13(old byte) byte {
	switch {
	case unicode.IsLower(rune(old)):
		return 'a' + (old-'a'+13)%26
	case unicode.IsUpper(rune(old)):
		return 'A' + (old-'A'+13)%26
	default:
		return old
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
