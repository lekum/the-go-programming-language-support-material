package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f, ok := w.(*os.File) // success
	if !ok {
		fmt.Fprintf(os.Stderr, "Interface %T is not of type *os.File\n", w)
	}
	b, ok := w.(*bytes.Buffer) // failure
	if !ok {
		fmt.Fprintf(os.Stderr, "Interface %T is not of type *bytes.Buffer\n", w)
	}
	fmt.Fprintf(os.Stdout, "%v => %T\n", f, f)
	fmt.Fprintf(os.Stdout, "%v => %T\n", b, b)
}
