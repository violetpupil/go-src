package runtime

func main() {
	if GOARCH != "wasm" {
		systemstack(func() {
			newm(sysmon, nil, -1)
		})
	}
}

func init() {
	go forcegchelper()
}

func forcegchelper() {
	for {
		gcStart(gcTrigger{kind: gcTriggerTime, now: nanotime()})
	}
}

func newm(fn func(), _p_ *p, id int64) {}

func sysmon() {}
