package runtime

func mallocgc() {
	shouldhelpgc := false

	if shouldhelpgc {
		if t := (gcTrigger{kind: gcTriggerHeap}); t.test() {
			gcStart(t)
		}
	}
}
