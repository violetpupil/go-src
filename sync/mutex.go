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

func (m *Mutex) lockSlow() {
	iter := 0
	old := m.state
	for {
		// locked && !starving && canSpin
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			runtime_doSpin()
			iter++
			continue
		}
		new := old
		// !starving
		if old&mutexStarving == 0 {
			new |= mutexLocked
		}
		// locked || starving
		if old&(mutexLocked|mutexStarving) != 0 {
			// waiter + 1
			new += 1 << mutexWaiterShift
		}
	}
}
