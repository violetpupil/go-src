package runtime

type mutex struct{}

type g struct{}

type m struct{}

type p struct{}

type forcegcstate struct {
	lock mutex
	g    *g
	idle uint32
}

var (
	forcegc forcegcstate
)
