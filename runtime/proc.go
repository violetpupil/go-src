package runtime

import "runtime/internal/atomic"

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
		if forcegc.idle != 0 {
			throw("forcegc: phase error")
		}
		atomic.Store(&forcegc.idle, 1)
		gcStart(gcTrigger{kind: gcTriggerTime, now: nanotime()})
	}
}

func newm(fn func(), _p_ *p, id int64) {}

func sysmon() {
	for {
		now := nanotime()

		if t := (gcTrigger{kind: gcTriggerTime, now: now}); t.test() && atomic.Load(&forcegc.idle) != 0 {
			forcegc.idle = 0
		}
	}
}
