package factorymethod

/*
製品を生成する工場を表現したインターフェイス
*/
type Factory interface {
	CreateProduct(string) Product
	RegisterProduct(Product)
}

type F struct {
	factory Factory
}

/*
CreateProductで製品を作り、
RegisterProductで登録する
という手順を定めた実装
*/
func (f *F) Create(owner string) Product {
	p := f.factory.CreateProduct(owner)
	f.factory.RegisterProduct(p)
	return p
}
