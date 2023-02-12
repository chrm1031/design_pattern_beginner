package factorymethod

func main() {
	factory := newIDCardFactory()
	f := &F{
		factory: factory,
	}
	card1 := f.Create("承太郎")
	card2 := f.Create("ポルナレフ")
	card3 := f.Create("ジョセフ")

	card1.use()
	card2.use()
	card3.use()
}
