package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// Fetch downloads the URL and returns the name
// and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

func main() {
	file, length, err := fetch(os.Args[1])
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	fmt.Printf("File name: %s\n", file)
	fmt.Printf("File length: %d bytes\n", length)
}
