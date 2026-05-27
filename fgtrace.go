package fgtrace

import (
	"io"
	"net/http"
	"time"

	"github.com/DataDog/gostackparse"
	"github.com/felixge/fgtrace/internal"
)

const (
	defaultFile         = "fgtrace.json"
	defaultHz           = 99
	defaultHTTPDuration = 30 * time.Second
	defaultStateFrames  = StateFramesRoot
)

// Config configures the capturing of traces as well as serving them via http.
// The zero value is a valid configuration.
type Config struct {
	// Hz determines how often the stack traces of all goroutines are captured
	// per second. WithDefaults() sets it to 99 Hz if it is 0.
	Hz int
	// IncludeSelf controls if the trace contains its own internal goroutines.
	// It's disabled by default because they are usually not of interest.
	IncludeSelf bool
	// StateFrames allows adding the state of goroutines as a virtual frame when
	// their stack traces are captured. WithDefaults() sets it to StateFramesRoot
	// if it is "".
	StateFrames StateFrames
	// Dst is the destination for traces created by calling Trace().
	// WithDefaults() sets it to File("fgtrace.json") if it is nil. Also see
	// Writer().
	Dst io.WriteCloser
	// HTTPDuration is the default duration for traces served via ServeHTTP().
	// WithDefaults() sets it to 30s if it is 0. It is ignored by Trace().
	HTTPDuration time.Duration
}

// StateFrames describes if and where virtual goroutine state frames are added.
type StateFrames string

const (
	// StateFramesRoot causes virtual goroutine state frames to be added at the
	// root of stack traces (e.g. above main).
	StateFramesRoot StateFrames = "root"
	// StateFramesLeaf causes virtual goroutine state frames to be added at the
	// leaf of stack traces.
	StateFramesLeaf StateFrames = "leaf"
	// StateFramesNo casuses no virtual goroutine state frames to be added to
	// stack traces.
	StateFramesNo StateFrames = "no"
)

// assert interface implementation
var _ http.Handler = Config{}

// WithDefaults returns a copy of c with default values applied as described in
// the type documentation. This is done automatically by Trace() and
// ServeHTTP(), but can be useful to log the effective configuration.
func (c Config) WithDefaults() Config { _ = "STUB: not implemented"; return *new(Config) }

// Trace applies WithDefaults to c and starts capturing a trace at c.Hz to
// c.Dst. Callers are responsible for calling Trace.Stop() to finish the trace.
func (c Config) Trace() *Trace { _ = "STUB: not implemented"; return nil }

// ServeHTTP applies WithDefaults to c and serves a trace. The query
// parameters "hz" and "seconds" can be used to overwrite the defaults.
func (c Config) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// File is a helper for Config.Dst that returns an io.WriteCloser that creates
// and writes to the file with the given name.
func File(name string) io.WriteCloser { _ = "STUB: not implemented"; return *new(io.WriteCloser) }

// Writer is a helper for for Config.Dst that returns an io.WriteCloser that
// writes to w and does nothing when Close() is called.
func Writer(w io.Writer) io.WriteCloser { _ = "STUB: not implemented"; return *new(io.WriteCloser) }

// Trace represents a trace that is being captured.
type Trace struct {
	c       Config            // config for the trace
	err     error             // error that caused the tracer to stop
	stop    chan struct{}     // closed to initiate stop
	stopped chan error        // messaged to confirm stop completed
	enc     *internal.Encoder // trace event format encoder
}

func (t *Trace) start() { _ = "STUB: not implemented"; return }

// Stop stops the trace, calls Close() on the configured dst and returns nil on
// success. Calling Stop() more than once returns the previous error or an
// error indicating that the tracer has already been stopped.
func (t *Trace) Stop() error { _ = "STUB: not implemented"; return nil }

// TODO(fg) does the trace format support writing error messages? if yes,
// we should probably attempt to write the error to the file as well.

// To be returned if Stop() is called more than once.

// trace is the background goroutine that takes goroutine profiles and
// converts them to trace events.
func (t *Trace) trace() error { _ = "STUB: not implemented"; return nil }

// Sleep until next tick comes up or the tracer is stopped.

type goroutineProfiler struct {
	buf []byte
}

func (g *goroutineProfiler) Goroutines() ([]*gostackparse.Goroutine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func excludeSelf(gs []*gostackparse.Goroutine) []*gostackparse.Goroutine {
	_ = "STUB: not implemented"
	return nil
}

func addVirualStateFrames(gs []*gostackparse.Goroutine, f StateFrames) {
	_ = "STUB: not implemented"
	return
}

// Taking a goroutine profile puts all running goroutines into runnable
// state. So let's indicate that we can't be sure of their real state,
// but that it's most likely running instead of runnable.
