package templatemethod

import "fmt"

/*
具象実装
*/
type CharDisplay struct {
	ch string
}

func (c *CharDisplay) Open() {
	fmt.Print("<<")
}
func (c *CharDisplay) Print() {
	fmt.Print(c.ch)
}
func (c *CharDisplay) Close() {
	fmt.Println(">>")
}
