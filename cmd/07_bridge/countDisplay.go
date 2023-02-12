package bridge

/*
機能のクラス階層

Displayに機能を追加したもの
Displayは「表示する」だけの機能だったが
CountDisplayは「指定回数だけ表示する」という機能を追加した
*/
type CountDisplay struct {
	Display
}

func newCountDisplay(i DisplayImpl) *CountDisplay {
	return &CountDisplay{
		Display{
			DisplayImpl: i,
		},
	}
}

// t回print()を繰り返す機能を追加
func (d *CountDisplay) multiDisplay(t int) {
	d.open()
	for i := 0; i < t; i++ {
		d.print()
	}
	d.close()
}
