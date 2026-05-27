package internal

import (
	"io"
	"os"
)

// NewFileWriter returns a io.WriteCloser.
func NewFileWriter(name string) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

type fileWriter struct {
	name       string
	firstWrite bool
	file       *os.File
	err        error
}

// Close implements io.Writer.
func (f *fileWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close implements io.Closer.
func (f *fileWriter) Close() error { _ = "STUB: not implemented"; return nil }

// String implements fmt.Stringer.
func (f *fileWriter) String() string {
	_ = "STUB: not implemented"

	// GoString implements fmt.GoStringer.
	return ""
}

func (f *fileWriter) GoString() string { _ = "STUB: not implemented"; return "" }
