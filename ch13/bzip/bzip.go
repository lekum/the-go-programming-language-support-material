// Package bzip provides a writer that uses bzip2 compression (bzip.org).
package bzip

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbz2
#include <bzlib.h>
int bz2compress(bz_stream *s, int action,
char *in, unsigned *inlen, char *out, unsigned *outlen);
*/
import "C"

import (
	"io"
	"unsafe"
)

type writer struct {
	w
	io.Writer // underlying output stream
	stream    *C.bz_stream
	outbuf    [64 * 1024]byte
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) io.WriteCloser {
	const (
		blockSize  = 9
		verbosity  = 0
		workFactor = 30
	)
	w := &writer{w: out, stream: new(C.bz_stream)}
	C.BZ2_bzCompressInit(w.stream, blockSize, verbosity, workFactor)
	return w
}
