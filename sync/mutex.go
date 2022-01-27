package sync

import "sync/atomic"

type Mutex struct {
	state int32
}

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

func (m *Mutex) Lock() {
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}
	m.lockSlow()
}

func (m *Mutex) lockSlow()
