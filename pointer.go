package main
import "fmt"
func change (a *int){
	*a = 100
}
func ret (b int) (a *(int)){
	a = &b
	return a
}
func main() {
	b := 25
	a := new(int)
	a = &b
	fmt.Printf("b co gia tri la %d va vi tri %v \n", *a, a)
	// thay doi gia tri cua b bang con tro 
	*a++
	fmt.Printf("gia tri b sau la: %d \n", b)
	// nhet con tro vao ham 
	change(a)
	fmt.Printf("gia tri luc nay cua b la %d \n", b)
	// ham trả về một con trỏ
	c:=100
	contro := ret(c)
	fmt.Println(contro)
}
// package main

// import (  
//     "fmt"
// )

// func modify(arr *[3]int) {  
//     (*arr)[0] = 90
// }

// func main() {  
//     a := [3]int{89, 90, 91}
//     modify(&a)
//     fmt.Println(a)
// }