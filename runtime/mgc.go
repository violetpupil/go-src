package runtime

type gcTrigger struct {
	kind gcTriggerKind
	now  int64
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
