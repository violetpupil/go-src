package runtime

type gcTrigger struct {
	kind gcTriggerKind
}

type gcTriggerKind int

const (
	gcTriggerHeap gcTriggerKind = iota
	gcTriggerTime
	gcTriggerCycle
)

func (t gcTrigger) test() bool {
	switch t.kind {
	case gcTriggerHeap:
	case gcTriggerTime:
	case gcTriggerCycle:
	}
	return true
}

func gcStart(trigger gcTrigger) {}
