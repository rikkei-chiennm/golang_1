package main

import (
	"fmt"
)
func main() {
	a:= [...]string{"messi","chien","thuy","truong","hai"}
	b:=a[2:]
	b[1]="ronaldo"
	fmt.Print(a)
	fmt.Println()
	fmt.Print(b)
	fmt.Print(c)
}
// i:=make([]int, 5, 5)
// khi slice cắt từ mảng được thay đổi ptu thì mảng cũng sẽ thay đổi theo, slice cũng thay đổi theo   
// cap trong slice được cắt từ 1 slice khác sẽ là từ phần tử đầu ( của slice mới )  tới cuối của slice cũ 