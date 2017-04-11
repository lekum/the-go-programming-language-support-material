package main

import (
	"fmt"
	"io"
	"os"
)

type CtWriter struct {
	Writer io.Writer
	Count  *int64
}

func (c CtWriter) Write(p []byte) (int, error) {
	n, err := c.Writer.Write(p)
	if err != nil {
		return n, err
	}
	*c.Count += int64(n)
	return int(*c.Count), err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := CtWriter{Writer: w, Count: new(int64)}
	return c, c.Count
}

func main() {
	f, err := os.Create("example.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	w, c := CountingWriter(f)
	fmt.Fprintf(w, "Message\n")
	fmt.Println("Count:", *c)
}
