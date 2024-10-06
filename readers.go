package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		for i := 0; i < n; i++ {
			fmt.Printf("b[i] = (%v, %v)\n", b[i], string(b[i]))
			// here is interesting: if slice rewrite data + n elem read => slice_len - n elems are garbage and just don't read (but exist).
		}
		fmt.Println()
		if err == io.EOF {
			break
		}
	}
}
