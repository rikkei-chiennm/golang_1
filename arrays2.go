package main
import "fmt"
func change(num [5]int) {
	num[0] = 55
	fmt.Println(num)
}
func main() {
	num:=[...]int{1,2,3,4,5}
	fmt.Println("mang ban dau", num)
	change(num)
	fmt.Println(num)
}