package main
import "fmt"
func main(){
// cdz:
// 	for i:=0;i<3;i++{
// 		for j:=1;j<4;j++{
// 			fmt.Printf("i = %d , j=%d\n", i, j)
// 			if i==j{
// 				break cdz
// 			}
// 		}
// 	}
a:=[...]int{1, 2, 3}
for _, v := range a { //ignores index  
	fmt.Print(v)
}
}