package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

/*
	func (r13 rot13Reader) Read() io.Reader {
		return r13.r
	}
*/
/*
func rot13(b byte) byte {
	switch {
	case b >= 'A' && b < 'Z'-13:
		return b + 13
	case b > 'Z'-13 && b <= 'Z':
		return b - 13
	case b >= 'a' && b < 'z' - 13:
		return b + 13
	case b > 'z' - 13 && b <= 'z':
		return b - 13
	default:
		panic(fmt.Sprintf("WARNING, %v, %v", b, string(b)))
	}
	return b
}
*/
func rot13(b byte) byte {
	switch {
	case b >= 'A' && b < 'Z'-13:
		b += 13
	case b > 'Z'-13 && b <= 'Z':
		b -= 13
	case b >= 'a' && b < 'z'-13:
		b += 13
	case b > 'z'-13 && b <= 'z':
		b -= 13
	}
	return b
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	r, err := r13.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := 0; i <= r; i++ {
		b[i] = rot13(b[i])
	}
	return r, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
