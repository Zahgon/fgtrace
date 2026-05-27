package internal

import "io"

var _ io.Writer = ErrWriter{}

type ErrWriter struct {
	Err error
}

func (e ErrWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
