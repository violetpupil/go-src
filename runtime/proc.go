package runtime

func init() {
	go forcegchelper()
}

func forcegchelper() {
	for {
		gcStart(gcTrigger{kind: gcTriggerTime, now: nanotime()})
	}
}
