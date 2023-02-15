package bridge

func main() {
	d1 := newDisplay(newStringDisplayImpl("Hello World"))
	d2 := newDisplay(newStringDisplayImpl(("Hello Japan")))
	d3 := newCountDisplay(newStringDisplayImpl(("Hello Universe")))

	// ランダム回数表示する機能を追加
	d4 := newRandomDisplay(newStringDisplayImpl(("Hello Random")))

	d1.display()
	d2.display()
	d3.display()
	d3.multiDisplay(5)
	// d3はdisplay()の他に追加されたメソッドmultiDisplay()も使用が可能

	d4.display()
	d4.randomDisplay()

	// テキストファイルの内容を表示する実装を追加
	const file = "hello.txt"
	d5 := newDisplay(newTextDisplayImpl(file))
	d5.display()

	d6 := newRandomDisplay(newTextDisplayImpl(file))
	d6.randomDisplay()

}
