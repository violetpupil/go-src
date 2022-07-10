package runtime

var gcController gcControllerState

type gcControllerState struct {
	trigger  uint64
	heapLive uint64
}
