package bridge

import (
	"math/rand"
	"time"
)

type RandomDisplay struct {
	Display
}

func newRandomDisplay(i DisplayImpl) *RandomDisplay {
	return &RandomDisplay{
		Display{
			DisplayImpl: i,
		},
	}
}

func (d *RandomDisplay) randomDisplay() {
	d.open()
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10)
	for i := 0; i < r; i++ {
		d.print()
	}
	d.close()
}
