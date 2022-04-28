package main
import "fmt"
func main(){
	a:=10
	b:=9
	if a%2==0 && b%2==0{
		fmt.Print("a,b la chan")
	}else{
		fmt.Print("a va b la khac nhau")
	}
}