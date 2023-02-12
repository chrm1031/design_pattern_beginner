package bridge

/*
実装のクラス階層
*/
type DisplayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}
