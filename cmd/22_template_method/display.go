package templatemethod

/*
テンプレートメソッド
*/
type IDisplay interface {
	Open()
	Print()
	Close()
}

type Display struct {
	iDisplay IDisplay
}

func (c *Display) display() {
	c.iDisplay.Open()
	for i := 0; i < 5; i++ {
		c.iDisplay.Print()
	}
	c.iDisplay.Close()
}
