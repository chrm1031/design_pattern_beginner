package templatemethod

import "fmt"

/*
具象実装
*/
type StringDisplay struct {
	text  string
	width int
}

func (s *StringDisplay) Open() {
	s.printLine()
}
func (s *StringDisplay) Print() {
	fmt.Printf("|%s|\n", s.text)
}
func (s *StringDisplay) Close() {
	s.printLine()
}

func (s *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
