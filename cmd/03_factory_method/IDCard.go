package factorymethod

import "fmt"

type IDCard struct {
	owner string
}

func new(o string) *IDCard {
	fmt.Printf("%sのカードを作ります。\n", o)
	return &IDCard{
		owner: o,
	}
}

func (c *IDCard) use() {
	fmt.Printf("%sのカードを使います。\n", c.owner)
}

func (c *IDCard) getOwner() string {
	return c.owner
}
