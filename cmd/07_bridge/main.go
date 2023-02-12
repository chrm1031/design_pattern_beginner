package bridge

func main() {
	d1 := newDisplay(newStringDisplayImpl("Hello World"))
	d2 := newDisplay(newStringDisplayImpl(("Hello Japan")))
	d3 := newCountDisplay(newStringDisplayImpl(("Hello Universe")))

	d1.display()
	d2.display()
	d3.display()
	d3.multiDisplay(5)
	// d3はdisplay()の他に追加されたメソッドmultiDisplay()も使用が可能
}
