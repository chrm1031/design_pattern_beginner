package bridge

/*
 */
type Display struct {
	DisplayImpl
}

func newDisplay(i DisplayImpl) *Display {
	return &Display{
		DisplayImpl: i,
	}
}

func (d *Display) open() {
	d.rawOpen()
}

func (d *Display) print() {
	d.rawPrint()
}

func (d *Display) close() {
	d.rawClose()
}

func (d *Display) display() {
	d.open()
	d.print()
	d.close()
}
