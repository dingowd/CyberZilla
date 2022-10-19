package lib1

import (
	"testcode/test2/lib2"
)

type Lib1 struct {
}

func New1() *Lib1 {
	return &Lib1{}
}

// Заменять возвращаемый тип, запрещено.
func (l Lib1) Do() lib2.SomeType {
	return lib2.SomeType{Data: "data from lib1"}
}

type Lib2 struct {
	l1 *Lib1
}

func New2(l1 *Lib1) *Lib2 {
	return &Lib2{
		l1: l1,
	}
}

func (l Lib2) Do() lib2.SomeType {
	return l.l1.Do()
}
