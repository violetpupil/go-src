package runtime

type mstats struct {
	enablegc         bool
	last_gc_nanotime uint64
}

var memstats mstats
