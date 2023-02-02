package factorymethod

/*
Template Method パターン
スーパークラスのほうにメソッド群が定義されている。
それらは抽象メソッドであり、具象な実装はサブクラスで行う
各サブクラスごとに実装が異なれば処理も異なる。

スーパークラスで処理の枠組みを決め、サブクラスで具体的実装を定義するデザインパターンを
Template Method パターンと呼ぶ
*/

func main() {
	_char := &CharDisplay{
		
	}
	charDisplay := &Display{
		iDisplay: _char,
	}
	charDisplay.display()
}
