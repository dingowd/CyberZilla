package main

import (
	"fmt"

	"testcode/test2/lib1"
)

// Вам необходимо решить проблему с "import cycle not allowed" с ограничениями, которые указаны в файлах lib1 и lib2.
// Если возникли сложности с данным заданием, пропускайте его и переходите к заданию test3.
func main() {
	l1 := lib1.New1()
	l2 := lib1.New2(l1)
	result := l2.Do()
	fmt.Println(result.Data)
}
