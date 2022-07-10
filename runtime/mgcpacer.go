package runtime

import "runtime/internal/atomic"

var gcController gcControllerState

type gcControllerState struct {
	gcPercent atomic.Int32
	trigger   uint64
	heapLive  uint64
}
