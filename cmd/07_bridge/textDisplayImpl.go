package bridge

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
実装のクラス階層

テキストファイルの内容を表示するクラス
*/
type TextDisplayImpl struct {
	fname string
	width int
	f     []byte
}

func newTextDisplayImpl(name string) *TextDisplayImpl {
	return &TextDisplayImpl{
		fname: name,
	}
}

func (d *TextDisplayImpl) rawOpen() {
	f, err := ioutil.ReadFile(d.fname)
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	d.width = len(f)
	d.f = f
	d.printline()
}

func (d *TextDisplayImpl) rawPrint() {
	fmt.Printf("|%s|\n", string(d.f))
}

func (d *TextDisplayImpl) rawClose() {
	d.printline()
}

func (d *TextDisplayImpl) printline() {
	fmt.Print("+")
	for i := 0; i < d.width; i++ {
		fmt.Print("-")
	}
	fmt.Print("+")
	fmt.Println()
}
