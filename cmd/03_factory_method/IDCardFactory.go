package factorymethod

type IDCardFactory struct {
	owners []string
}

func newIDCardFactory() Factory {
	return &IDCardFactory{}
}

func (f *IDCardFactory) CreateProduct(o string) Product {
	return new(o)
}

func (f *IDCardFactory) RegisterProduct(p Product) {
	f.owners = append(f.owners, p.(*IDCard).getOwner())
}
