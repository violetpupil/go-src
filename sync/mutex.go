package sync

import "sync/atomic"

func throw(string)

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

func (m *Mutex) lockSlow() {
	old := m.state
	for {
		new := old
		// !starving
		if old&mutexStarving == 0 {
			new |= mutexLocked
		}
		// 101 || 1 || 100
		if old&(mutexLocked|mutexStarving) != 0 {
			new += 1 << mutexWaiterShift
		}
	}
}
