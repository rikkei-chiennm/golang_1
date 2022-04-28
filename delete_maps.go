package main
import "fmt"
func main() {
	bando := map[string]int{
		"vietnam" : 90,
		"USA" : 900,
		"thailand": 65,
	}
	fmt.Print(bando)
	delete(bando,"vietnam")
	fmt.Println()
	fmt.Print(bando)
 }