package lib2

/*type Lib2 struct {
	l1 *lib1.Lib1
}*/

// Перемещать объявление данного типа запрещено.
type SomeType struct {
	Data string
}

/*func New(l1 *lib1.Lib1) *Lib2 {
	return &Lib2{
		l1: l1,
	}
}

func (l Lib2) Do() SomeType {
	return l.l1.Do()
}*/
