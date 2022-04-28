package main
import "fmt"
func main() {
	sum:=0.00
	a:=[...]float64{1.2,2,3,4.5,5,6.2}
	for i,v := range a{
		fmt.Printf("a[%d] co gia tri la: %.2f \n", i,v)
		sum+=v
	}
	fmt.Println(sum)
}