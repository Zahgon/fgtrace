package internal

// Implementation of Chrome's Trace Event Format, see:
// https://docs.google.com/document/d/1CvAClvFfyA5R-PhYUmn5OOQtYMH4h6I0nSsKchNAySU/preview

import (
	"encoding/json"
	"io"

	"github.com/DataDog/gostackparse"
)

func NewEncoder(w io.Writer) (*Encoder, error) { _ = "STUB: not implemented"; return nil, nil }

func Unmarshal(data []byte) (*TraceData, error) { _ = "STUB: not implemented"; return nil, nil }

type TraceData struct {
	Events []*Event
}

func (t *TraceData) MetaHz() int { _ = "STUB: not implemented"; return 0 }

func (t *TraceData) Filter(fn func(*Event) bool) *TraceData { _ = "STUB: not implemented"; return nil }

// CallGraph returns a graph of all function calls (stack traces) in the trace.
func (t *TraceData) CallGraph() *Node { _ = "STUB: not implemented"; return nil }

func (t *TraceData) Len() int { _ = "STUB: not implemented"; return 0 }

type Node struct {
	Func     string
	Children []*Node
}

func (n *Node) HasLeaf(fn string) bool { _ = "STUB: not implemented"; return false }

func (n *Node) String() string { _ = "STUB: not implemented"; return "" }

type Event struct {
	Name string `json:"name,omitempty"`
	Ph   string `json:"ph,omitempty"`
	// Ts is the tracing clock timestamp of the event. The timestamps are
	// provided at microsecond granularity.
	Ts   float64                `json:"ts"`
	Pid  int64                  `json:"pid,omitempty"`
	Tid  int64                  `json:"tid,omitempty"`
	Args map[string]interface{} `json:"args,omitempty"`
}

// Encoder implements a small subset of the "Trace Event Format" spec needed to
// make fgtrace output data that can be displayed by perfetto.dev.
// https://docs.google.com/document/d/1CvAClvFfyA5R-PhYUmn5OOQtYMH4h6I0nSsKchNAySU/preview
type Encoder struct {
	w     io.Writer
	json  json.Encoder
	first bool
}

func (e *Encoder) CustomMeta(name string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) Encode(ts float64, prev, current *gostackparse.Goroutine) error {
	_ = "STUB: not implemented"
	return nil
}

// Determine the number of stack frames that are identical between prev and
// current going from root frame (e.g. main) to the leaf frame.

// Emit end events for prev stack frames that are no longer part of the
// current stack going from leaf to root frame.

// Emit start events for current stack frames that were not part of the prev
// stack going from root to leaf frame.

func (e *Encoder) encode(ev *Event) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) Finish() error { _ = "STUB: not implemented"; return nil }
