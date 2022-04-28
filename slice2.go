package main
import "fmt"
func main() {
	// cars := []string{"toyota", "mers", "bmw", "honda"}
	// cars = append(cars, "huyndaii")
	// fmt.Println(cars, " ", len(cars), " ", cap(cars))
	var names []string
	if names==nil{
		fmt.Println("slice is nil going to apppend")
		names = append(names, "hay", " that", " co", " gang")
		fmt.Print(names)
	}
}
