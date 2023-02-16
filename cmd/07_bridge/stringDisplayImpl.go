package bridge

import "fmt"

/*
実装のクラス階層

文字列を表示するクラス
*/
type StringDisplayImpl struct {
	text  string
	width int
}

func newStringDisplayImpl(t string) *StringDisplayImpl {
	return &StringDisplayImpl{
		text:  t,
		width: len(t),
	}
}

func (d *StringDisplayImpl) rawOpen() {
	d.printline()
}

func (d *StringDisplayImpl) rawPrint() {
	fmt.Printf("|%s|\n", d.text)
}

func (d *StringDisplayImpl) rawClose() {
	d.printline()
}

func (d *StringDisplayImpl) printline() {
	fmt.Print("+")
	for i := 0; i < d.width; i++ {
		fmt.Print("-")
	}
	fmt.Print("+")
	fmt.Println()
}
