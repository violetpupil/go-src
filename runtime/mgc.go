package runtime

import (
	"runtime/internal/atomic"
)

var gcphase uint32

const (
	_GCoff = iota
	_GCmark
	_GCmarktermination
)

var work struct {
	cycles uint32
}

func GC() {
	n := atomic.Load(&work.cycles)

	gcStart(gcTrigger{kind: gcTriggerCycle, n: n + 1})
}

type gcTrigger struct {
	kind gcTriggerKind
	now  int64
	n    uint32
}

type gcTriggerKind int

const (
	gcTriggerHeap gcTriggerKind = iota
	gcTriggerTime
	gcTriggerCycle
)

func (t gcTrigger) test() bool {
	if !memstats.enablegc || panicking != 0 || gcphase != _GCoff {
		return false
	}
	switch t.kind {
	case gcTriggerHeap:
		return gcController.heapLive >= gcController.trigger
	case gcTriggerTime:
	case gcTriggerCycle:
	}
	return true
}

func gcStart(trigger gcTrigger) {}
