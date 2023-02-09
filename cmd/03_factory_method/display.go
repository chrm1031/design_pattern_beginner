package factorymethod

/*
テンプレートメソッド
*/
type IDisplay interface {
	open()
	print()
	close()
}

type Display struct {
	iDisplay IDisplay
}

func (c *Display) display() {
	c.iDisplay.open()
	for i := 0; i < 5; i++ {
		c.iDisplay.print()
	}
	c.iDisplay.close()
}
